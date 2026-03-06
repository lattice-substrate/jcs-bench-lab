use crate::jcserr::{FailureClass, JcsError};

pub const DEFAULT_MAX_DEPTH: usize = 1000;
pub const DEFAULT_MAX_INPUT_SIZE: usize = 64 * 1024 * 1024;
pub const DEFAULT_MAX_VALUES: usize = 1_000_000;
pub const DEFAULT_MAX_OBJECT_MEMBERS: usize = 250_000;
pub const DEFAULT_MAX_ARRAY_ELEMENTS: usize = 250_000;
pub const DEFAULT_MAX_STRING_BYTES: usize = 8 * 1024 * 1024;
pub const DEFAULT_MAX_NUMBER_CHARS: usize = 4096;

#[derive(Debug, Clone, PartialEq)]
pub enum Value {
    Null,
    Bool(String),
    Number(f64),
    String(String),
    Array(Vec<Value>),
    Object(Vec<Member>),
}

#[derive(Debug, Clone, PartialEq)]
pub struct Member {
    pub key: String,
    pub value: Value,
}

#[derive(Debug, Clone)]
pub struct Options {
    pub max_depth: usize,
    pub max_input_size: usize,
    pub max_values: usize,
    pub max_object_members: usize,
    pub max_array_elements: usize,
    pub max_string_bytes: usize,
    pub max_number_chars: usize,
}

impl Default for Options {
    fn default() -> Self {
        Self {
            max_depth: DEFAULT_MAX_DEPTH,
            max_input_size: DEFAULT_MAX_INPUT_SIZE,
            max_values: DEFAULT_MAX_VALUES,
            max_object_members: DEFAULT_MAX_OBJECT_MEMBERS,
            max_array_elements: DEFAULT_MAX_ARRAY_ELEMENTS,
            max_string_bytes: DEFAULT_MAX_STRING_BYTES,
            max_number_chars: DEFAULT_MAX_NUMBER_CHARS,
        }
    }
}

pub fn parse(data: &[u8]) -> Result<Value, JcsError> {
    parse_with_options(data, &Options::default())
}

pub fn parse_with_options(data: &[u8], opts: &Options) -> Result<Value, JcsError> {
    if data.len() > opts.max_input_size {
        return Err(JcsError::new(
            FailureClass::BoundExceeded,
            0,
            format!(
                "input size {} exceeds maximum {}",
                data.len(),
                opts.max_input_size
            ),
        ));
    }

    if std::str::from_utf8(data).is_err() {
        let off = first_invalid_utf8_offset(data);
        return Err(JcsError::new(
            FailureClass::InvalidUtf8,
            off as i64,
            "input is not valid UTF-8",
        ));
    }

    let mut p = Parser {
        data,
        pos: 0,
        depth: 0,
        value_count: 0,
        opts,
    };

    p.skip_whitespace();
    let v = p.parse_value()?;
    p.skip_whitespace();
    if p.pos != p.data.len() {
        return Err(p.new_error("trailing content after JSON value"));
    }
    Ok(v)
}

fn first_invalid_utf8_offset(data: &[u8]) -> usize {
    let mut i = 0;
    while i < data.len() {
        let b = data[i];
        let seq_len = if b < 0x80 {
            1
        } else if b < 0xC0 {
            return i;
        } else if b < 0xE0 {
            2
        } else if b < 0xF0 {
            3
        } else if b < 0xF8 {
            4
        } else {
            return i;
        };
        if i + seq_len > data.len() {
            return i;
        }
        if std::str::from_utf8(&data[i..i + seq_len]).is_err() {
            return i;
        }
        i += seq_len;
    }
    0
}

struct Parser<'a> {
    data: &'a [u8],
    pos: usize,
    depth: usize,
    value_count: usize,
    opts: &'a Options,
}

impl<'a> Parser<'a> {
    fn new_error(&self, msg: &str) -> JcsError {
        JcsError::new(FailureClass::InvalidGrammar, self.pos as i64, msg)
    }

    fn new_class_error(&self, class: FailureClass, msg: String) -> JcsError {
        JcsError::new(class, self.pos as i64, msg)
    }

    fn peek(&self) -> Option<u8> {
        if self.pos < self.data.len() {
            Some(self.data[self.pos])
        } else {
            None
        }
    }

    fn next_byte(&mut self) -> Option<u8> {
        if self.pos < self.data.len() {
            let b = self.data[self.pos];
            self.pos += 1;
            Some(b)
        } else {
            None
        }
    }

    fn expect(&mut self, expected: u8) -> Result<(), JcsError> {
        match self.next_byte() {
            None => Err(JcsError::new(
                FailureClass::InvalidGrammar,
                self.pos as i64,
                format!(
                    "unexpected end of input, expected {:?}",
                    expected as char
                ),
            )),
            Some(b) if b != expected => Err(JcsError::new(
                FailureClass::InvalidGrammar,
                (self.pos - 1) as i64,
                format!("expected {:?}, got {:?}", expected as char, b as char),
            )),
            _ => Ok(()),
        }
    }

    fn skip_whitespace(&mut self) {
        while self.pos < self.data.len() {
            match self.data[self.pos] {
                b' ' | b'\t' | b'\n' | b'\r' => self.pos += 1,
                _ => return,
            }
        }
    }

    fn push_depth(&mut self) -> Result<(), JcsError> {
        self.depth += 1;
        if self.depth > self.opts.max_depth {
            return Err(self.new_class_error(
                FailureClass::BoundExceeded,
                format!(
                    "nesting depth {} exceeds maximum {}",
                    self.depth, self.opts.max_depth
                ),
            ));
        }
        Ok(())
    }

    fn pop_depth(&mut self) {
        self.depth -= 1;
    }

    fn parse_value(&mut self) -> Result<Value, JcsError> {
        self.value_count += 1;
        if self.value_count > self.opts.max_values {
            return Err(self.new_class_error(
                FailureClass::BoundExceeded,
                format!(
                    "value count {} exceeds maximum {}",
                    self.value_count, self.opts.max_values
                ),
            ));
        }

        match self.peek() {
            None => Err(self.new_error("unexpected end of input")),
            Some(b'{') => self.parse_object(),
            Some(b'[') => self.parse_array(),
            Some(b'"') => self.parse_string_value(),
            Some(b't') | Some(b'f') => self.parse_bool(),
            Some(b'n') => self.parse_null(),
            Some(_) => self.parse_number(),
        }
    }

    fn parse_object(&mut self) -> Result<Value, JcsError> {
        self.push_depth()?;
        self.expect(b'{')?;
        self.skip_whitespace();

        let mut members = Vec::new();
        let mut seen = std::collections::HashMap::new();

        match self.peek() {
            None => {
                self.pop_depth();
                return Err(self.new_error("unexpected end of input in object"));
            }
            Some(b'}') => {
                self.pos += 1;
                self.pop_depth();
                return Ok(Value::Object(members));
            }
            _ => {}
        }

        loop {
            self.skip_whitespace();
            let key_start = self.pos;
            let key_val = self.parse_string_value()?;
            let key = match key_val {
                Value::String(s) => s,
                _ => unreachable!(),
            };

            if let Some(&first_off) = seen.get(&key) {
                self.pop_depth();
                return Err(JcsError::new(
                    FailureClass::DuplicateKey,
                    key_start as i64,
                    format!(
                        "duplicate object key {:?} (first at byte {})",
                        key, first_off
                    ),
                ));
            }
            seen.insert(key.clone(), key_start);

            self.skip_whitespace();
            self.expect(b':')?;
            self.skip_whitespace();

            let val = self.parse_value()?;

            if members.len() >= self.opts.max_object_members {
                self.pop_depth();
                return Err(self.new_class_error(
                    FailureClass::BoundExceeded,
                    format!(
                        "object member count exceeds maximum {}",
                        self.opts.max_object_members
                    ),
                ));
            }
            members.push(Member { key, value: val });

            self.skip_whitespace();
            match self.peek() {
                None => {
                    self.pop_depth();
                    return Err(self.new_error("unexpected end of input in object"));
                }
                Some(b'}') => {
                    self.pos += 1;
                    self.pop_depth();
                    return Ok(Value::Object(members));
                }
                Some(b',') => {
                    self.pos += 1;
                    continue;
                }
                Some(c) => {
                    self.pop_depth();
                    return Err(self.new_error(&format!(
                        "expected ',' or '}}' in object, got {:?}",
                        c as char
                    )));
                }
            }
        }
    }

    fn parse_array(&mut self) -> Result<Value, JcsError> {
        self.push_depth()?;
        self.expect(b'[')?;
        self.skip_whitespace();

        let mut elems = Vec::new();

        match self.peek() {
            None => {
                self.pop_depth();
                return Err(self.new_error("unexpected end of input in array"));
            }
            Some(b']') => {
                self.pos += 1;
                self.pop_depth();
                return Ok(Value::Array(elems));
            }
            _ => {}
        }

        loop {
            self.skip_whitespace();
            let elem = self.parse_value()?;

            if elems.len() >= self.opts.max_array_elements {
                self.pop_depth();
                return Err(self.new_class_error(
                    FailureClass::BoundExceeded,
                    format!(
                        "array element count exceeds maximum {}",
                        self.opts.max_array_elements
                    ),
                ));
            }
            elems.push(elem);

            self.skip_whitespace();
            match self.peek() {
                None => {
                    self.pop_depth();
                    return Err(self.new_error("unexpected end of input in array"));
                }
                Some(b']') => {
                    self.pos += 1;
                    self.pop_depth();
                    return Ok(Value::Array(elems));
                }
                Some(b',') => {
                    self.pos += 1;
                    continue;
                }
                Some(c) => {
                    self.pop_depth();
                    return Err(self.new_error(&format!(
                        "expected ',' or ']' in array, got {:?}",
                        c as char
                    )));
                }
            }
        }
    }

    fn parse_string_value(&mut self) -> Result<Value, JcsError> {
        let s = self.parse_string_inner()?;
        Ok(Value::String(s))
    }

    fn parse_string_inner(&mut self) -> Result<String, JcsError> {
        self.expect(b'"')?;

        // Fast path: scan for closing quote over pure printable ASCII
        let start = self.pos;
        while self.pos < self.data.len() {
            let b = self.data[self.pos];
            if b == b'"' {
                let s = &self.data[start..self.pos];
                if s.len() > self.opts.max_string_bytes {
                    return Err(self.new_class_error(
                        FailureClass::BoundExceeded,
                        format!(
                            "string decoded length exceeds maximum {} bytes",
                            self.opts.max_string_bytes
                        ),
                    ));
                }
                self.pos += 1;
                // SAFETY: we verified UTF-8 for the whole input at parse entry
                return Ok(std::str::from_utf8(s).unwrap().to_string());
            }
            if b < 0x20 || b == b'\\' || b >= 0x80 {
                break;
            }
            self.pos += 1;
        }

        // General path with escape/non-ASCII handling
        let mut buf = Vec::with_capacity(self.pos - start + 32);
        buf.extend_from_slice(&self.data[start..self.pos]);

        loop {
            if self.pos >= self.data.len() {
                return Err(self.new_error("unterminated string"));
            }
            let b = self.data[self.pos];
            if b == b'"' {
                self.pos += 1;
                // SAFETY: we build valid UTF-8 from validated input
                return Ok(String::from_utf8(buf).unwrap());
            }
            if b == b'\\' {
                let escape_start = self.pos;
                self.pos += 1;
                let r = self.parse_escape(escape_start)?;
                validate_string_rune(r, escape_start)?;
                let mut tmp = [0u8; 4];
                let encoded = r.encode_utf8(&mut tmp);
                if buf.len() + encoded.len() > self.opts.max_string_bytes {
                    return Err(self.new_class_error(
                        FailureClass::BoundExceeded,
                        format!(
                            "string decoded length exceeds maximum {} bytes",
                            self.opts.max_string_bytes
                        ),
                    ));
                }
                buf.extend_from_slice(encoded.as_bytes());
                continue;
            }
            if b < 0x20 {
                return Err(JcsError::new(
                    FailureClass::InvalidGrammar,
                    self.pos as i64,
                    format!("unescaped control character 0x{:02X} in string", b),
                ));
            }
            // Copy UTF-8 character
            let source_offset = self.pos;
            let remaining = &self.data[self.pos..];
            let s = match std::str::from_utf8(remaining) {
                Ok(s) => s,
                Err(e) => {
                    let valid_up_to = e.valid_up_to();
                    if valid_up_to == 0 {
                        return Err(JcsError::new(
                            FailureClass::InvalidUtf8,
                            self.pos as i64,
                            format!("invalid UTF-8 byte 0x{:02X} in string", b),
                        ));
                    }
                    // SAFETY: valid_up_to bytes are valid UTF-8
                    unsafe { std::str::from_utf8_unchecked(&remaining[..valid_up_to]) }
                }
            };
            let ch = s.chars().next().unwrap();
            let size = ch.len_utf8();
            validate_string_rune(ch, source_offset)?;
            if buf.len() + size > self.opts.max_string_bytes {
                return Err(self.new_class_error(
                    FailureClass::BoundExceeded,
                    format!(
                        "string decoded length exceeds maximum {} bytes",
                        self.opts.max_string_bytes
                    ),
                ));
            }
            buf.extend_from_slice(&self.data[self.pos..self.pos + size]);
            self.pos += size;
        }
    }

    fn parse_escape(&mut self, source_offset: usize) -> Result<char, JcsError> {
        if self.pos >= self.data.len() {
            return Err(JcsError::new(
                FailureClass::InvalidGrammar,
                source_offset as i64,
                "unterminated escape sequence",
            ));
        }
        let b = self.data[self.pos];
        self.pos += 1;

        if b == b'u' {
            return self.parse_unicode_escape(source_offset);
        }
        match escaped_rune(b) {
            Some(r) => Ok(r),
            None => Err(JcsError::new(
                FailureClass::InvalidGrammar,
                source_offset as i64,
                format!("invalid escape character {:?}", b as char),
            )),
        }
    }

    fn parse_unicode_escape(&mut self, source_offset: usize) -> Result<char, JcsError> {
        let r1 = self.read_hex4(source_offset)?;

        if !(0xD800..=0xDFFF).contains(&r1) {
            return Ok(char::from_u32(r1).unwrap());
        }

        // Lone low surrogate
        if r1 >= 0xDC00 {
            return Err(JcsError::new(
                FailureClass::LoneSurrogate,
                source_offset as i64,
                format!("lone low surrogate U+{:04X}", r1),
            ));
        }

        // High surrogate must be followed by \uXXXX low surrogate
        if self.pos + 1 >= self.data.len()
            || self.data[self.pos] != b'\\'
            || self.data[self.pos + 1] != b'u'
        {
            return Err(JcsError::new(
                FailureClass::LoneSurrogate,
                source_offset as i64,
                format!("lone high surrogate U+{:04X} (no following \\u)", r1),
            ));
        }
        let second_escape_offset = self.pos;
        self.pos += 2;

        let r2 = self.read_hex4(second_escape_offset)?;
        if !(0xDC00..=0xDFFF).contains(&r2) {
            return Err(JcsError::new(
                FailureClass::LoneSurrogate,
                second_escape_offset as i64,
                format!(
                    "high surrogate U+{:04X} followed by non-low-surrogate U+{:04X}",
                    r1, r2
                ),
            ));
        }

        // Decode surrogate pair
        let decoded = 0x10000 + ((r1 - 0xD800) << 10) + (r2 - 0xDC00);
        match char::from_u32(decoded) {
            Some(c) => Ok(c),
            None => Err(JcsError::new(
                FailureClass::LoneSurrogate,
                source_offset as i64,
                format!("invalid surrogate pair U+{:04X} U+{:04X}", r1, r2),
            )),
        }
    }

    fn read_hex4(&mut self, source_offset: usize) -> Result<u32, JcsError> {
        if self.pos + 4 > self.data.len() {
            return Err(JcsError::new(
                FailureClass::InvalidGrammar,
                source_offset as i64,
                "incomplete \\u escape",
            ));
        }
        let mut val = 0u32;
        for i in 0..4 {
            match hex_val(self.data[self.pos + i]) {
                Some(d) => val = (val << 4) | d as u32,
                None => {
                    let hex =
                        std::str::from_utf8(&self.data[self.pos..self.pos + 4]).unwrap_or("????");
                    return Err(JcsError::new(
                        FailureClass::InvalidGrammar,
                        source_offset as i64,
                        format!("invalid hex in \\u escape: {:?}", hex),
                    ));
                }
            }
        }
        self.pos += 4;
        Ok(val)
    }

    fn parse_number(&mut self) -> Result<Value, JcsError> {
        let start = self.pos;

        // Optional minus sign
        if self.pos < self.data.len() && self.data[self.pos] == b'-' {
            self.pos += 1;
        }

        // Integer part
        self.scan_integer_part(start)?;
        // Optional fraction
        self.scan_fraction_part(start)?;
        // Optional exponent
        self.scan_exponent_part(start)?;

        if self.pos - start > self.opts.max_number_chars {
            return Err(JcsError::new(
                FailureClass::BoundExceeded,
                start as i64,
                format!(
                    "number token length {} exceeds maximum {}",
                    self.pos - start,
                    self.opts.max_number_chars
                ),
            ));
        }

        let raw = std::str::from_utf8(&self.data[start..self.pos]).unwrap();
        self.build_number_value(start, raw)
    }

    fn scan_integer_part(&mut self, num_start: usize) -> Result<(), JcsError> {
        if self.pos >= self.data.len() {
            return Err(self.new_error("unexpected end of input in number"));
        }
        if self.data[self.pos] == b'0' {
            self.pos += 1;
            if self.pos < self.data.len() && is_digit(self.data[self.pos]) {
                return Err(self.new_error("leading zero in number"));
            }
            return Ok(());
        }
        if self.data[self.pos] < b'1' || self.data[self.pos] > b'9' {
            return Err(self.new_error(&format!(
                "invalid number character {:?}",
                self.data[self.pos] as char
            )));
        }
        while self.pos < self.data.len() && is_digit(self.data[self.pos]) {
            self.pos += 1;
            if self.pos - num_start > self.opts.max_number_chars {
                return Err(JcsError::new(
                    FailureClass::BoundExceeded,
                    num_start as i64,
                    format!(
                        "number token length {} exceeds maximum {}",
                        self.pos - num_start,
                        self.opts.max_number_chars
                    ),
                ));
            }
        }
        Ok(())
    }

    fn scan_fraction_part(&mut self, num_start: usize) -> Result<(), JcsError> {
        if self.pos >= self.data.len() || self.data[self.pos] != b'.' {
            return Ok(());
        }
        self.pos += 1;
        if self.pos >= self.data.len() || !is_digit(self.data[self.pos]) {
            return Err(self.new_error("expected digit after decimal point"));
        }
        while self.pos < self.data.len() && is_digit(self.data[self.pos]) {
            self.pos += 1;
            if self.pos - num_start > self.opts.max_number_chars {
                return Err(JcsError::new(
                    FailureClass::BoundExceeded,
                    num_start as i64,
                    format!(
                        "number token length {} exceeds maximum {}",
                        self.pos - num_start,
                        self.opts.max_number_chars
                    ),
                ));
            }
        }
        Ok(())
    }

    fn scan_exponent_part(&mut self, num_start: usize) -> Result<(), JcsError> {
        if self.pos >= self.data.len()
            || (self.data[self.pos] != b'e' && self.data[self.pos] != b'E')
        {
            return Ok(());
        }
        self.pos += 1;
        if self.pos < self.data.len()
            && (self.data[self.pos] == b'+' || self.data[self.pos] == b'-')
        {
            self.pos += 1;
        }
        if self.pos >= self.data.len() || !is_digit(self.data[self.pos]) {
            return Err(self.new_error("expected digit in exponent"));
        }
        while self.pos < self.data.len() && is_digit(self.data[self.pos]) {
            self.pos += 1;
            if self.pos - num_start > self.opts.max_number_chars {
                return Err(JcsError::new(
                    FailureClass::BoundExceeded,
                    num_start as i64,
                    format!(
                        "number token length {} exceeds maximum {}",
                        self.pos - num_start,
                        self.opts.max_number_chars
                    ),
                ));
            }
        }
        Ok(())
    }

    fn build_number_value(&self, start: usize, raw: &str) -> Result<Value, JcsError> {
        let f: f64 = match raw.parse() {
            Ok(v) => v,
            Err(e) => {
                // Check if it's a range error (overflow)
                let f_check: Result<f64, _> = raw.parse();
                if let Ok(v) = f_check {
                    if v.is_infinite() {
                        return Err(JcsError::new(
                            FailureClass::NumberOverflow,
                            start as i64,
                            "number overflows IEEE 754 double",
                        ));
                    }
                }
                return Err(JcsError::new(
                    FailureClass::InvalidGrammar,
                    start as i64,
                    format!("invalid number: {}", e),
                ));
            }
        };

        if f.is_nan() || f.is_infinite() {
            return Err(JcsError::new(
                FailureClass::NumberOverflow,
                start as i64,
                "number overflows IEEE 754 double",
            ));
        }

        // Lexical negative zero
        if raw.starts_with('-') && token_represents_zero(raw) {
            return Err(JcsError::new(
                FailureClass::NumberNegZero,
                start as i64,
                "negative zero token is not allowed",
            ));
        }

        // Non-zero underflows to zero
        if f == 0.0 && !token_represents_zero(raw) {
            return Err(JcsError::new(
                FailureClass::NumberUnderflow,
                start as i64,
                "non-zero number underflows to IEEE 754 zero",
            ));
        }

        Ok(Value::Number(f))
    }

    fn parse_bool(&mut self) -> Result<Value, JcsError> {
        if self.pos + 4 <= self.data.len() && &self.data[self.pos..self.pos + 4] == b"true" {
            self.pos += 4;
            return Ok(Value::Bool("true".to_string()));
        }
        if self.pos + 5 <= self.data.len() && &self.data[self.pos..self.pos + 5] == b"false" {
            self.pos += 5;
            return Ok(Value::Bool("false".to_string()));
        }
        Err(self.new_error("invalid literal"))
    }

    fn parse_null(&mut self) -> Result<Value, JcsError> {
        if self.pos + 4 <= self.data.len() && &self.data[self.pos..self.pos + 4] == b"null" {
            self.pos += 4;
            return Ok(Value::Null);
        }
        Err(self.new_error("invalid literal"))
    }
}

fn escaped_rune(b: u8) -> Option<char> {
    match b {
        b'"' => Some('"'),
        b'\\' => Some('\\'),
        b'/' => Some('/'),
        b'b' => Some('\u{0008}'),
        b'f' => Some('\u{000C}'),
        b'n' => Some('\n'),
        b'r' => Some('\r'),
        b't' => Some('\t'),
        _ => None,
    }
}

fn hex_val(b: u8) -> Option<u8> {
    match b {
        b'0'..=b'9' => Some(b - b'0'),
        b'a'..=b'f' => Some(b - b'a' + 10),
        b'A'..=b'F' => Some(b - b'A' + 10),
        _ => None,
    }
}

fn validate_string_rune(r: char, source_offset: usize) -> Result<(), JcsError> {
    if is_noncharacter(r) {
        return Err(JcsError::new(
            FailureClass::Noncharacter,
            source_offset as i64,
            format!("string contains Unicode noncharacter U+{:04X}", r as u32),
        ));
    }
    let cp = r as u32;
    if (0xD800..=0xDFFF).contains(&cp) {
        return Err(JcsError::new(
            FailureClass::LoneSurrogate,
            source_offset as i64,
            format!("string contains surrogate code point U+{:04X}", cp),
        ));
    }
    Ok(())
}

pub fn is_noncharacter(r: char) -> bool {
    let cp = r as u32;
    if (0xFDD0..=0xFDEF).contains(&cp) {
        return true;
    }
    if cp & 0xFFFE == 0xFFFE && cp <= 0x10FFFF {
        return true;
    }
    false
}

fn token_represents_zero(raw: &str) -> bool {
    let bytes = raw.as_bytes();
    let start = if !bytes.is_empty() && bytes[0] == b'-' {
        1
    } else {
        0
    };
    let end = bytes
        .iter()
        .position(|&b| b == b'e' || b == b'E')
        .unwrap_or(bytes.len());
    for &b in &bytes[start..end] {
        if (b'1'..=b'9').contains(&b) {
            return false;
        }
    }
    true
}

fn is_digit(b: u8) -> bool {
    b.is_ascii_digit()
}
