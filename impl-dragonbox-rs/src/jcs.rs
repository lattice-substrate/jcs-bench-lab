use crate::jcserr::{FailureClass, JcsError};
use crate::jcsfloat;
use crate::jcstoken::{self, Member, Value};

/// Canonicalize parses JSON input and produces the RFC 8785 JCS canonical byte
/// sequence. Equivalent to parse followed by serialize.
pub fn canonicalize(input: &[u8]) -> Result<Vec<u8>, JcsError> {
    let v = jcstoken::parse(input)?;
    let mut buf = Vec::with_capacity(input.len());
    serialize_value(&mut buf, &v)?;
    Ok(buf)
}

/// Verify checks if input is already in canonical form by parsing, serializing,
/// and comparing bytes.
pub fn verify(input: &[u8]) -> Result<(), JcsError> {
    let v = jcstoken::parse(input)?;
    let mut buf = Vec::with_capacity(input.len());
    serialize_value(&mut buf, &v)?;
    if input != buf.as_slice() {
        return Err(JcsError::new(
            FailureClass::NotCanonical,
            -1,
            "input is not canonical",
        ));
    }
    Ok(())
}

fn serialize_value(buf: &mut Vec<u8>, v: &Value) -> Result<(), JcsError> {
    match v {
        Value::Null => {
            buf.extend_from_slice(b"null");
            Ok(())
        }
        Value::Bool(s) => {
            buf.extend_from_slice(s.as_bytes());
            Ok(())
        }
        Value::Number(f) => {
            let s = jcsfloat::format_double(*f)?;
            buf.extend_from_slice(s.as_bytes());
            Ok(())
        }
        Value::String(s) => {
            serialize_string(buf, s);
            Ok(())
        }
        Value::Array(elems) => {
            buf.push(b'[');
            for (i, elem) in elems.iter().enumerate() {
                if i > 0 {
                    buf.push(b',');
                }
                serialize_value(buf, elem)?;
            }
            buf.push(b']');
            Ok(())
        }
        Value::Object(members) => serialize_object(buf, members),
    }
}

/// RFC 8785 string escaping rules (§3.2.2.2).
fn serialize_string(buf: &mut Vec<u8>, s: &str) {
    buf.push(b'"');
    for &b in s.as_bytes() {
        match b {
            b'"' => buf.extend_from_slice(b"\\\""),
            b'\\' => buf.extend_from_slice(b"\\\\"),
            0x08 => buf.extend_from_slice(b"\\b"),
            0x09 => buf.extend_from_slice(b"\\t"),
            0x0A => buf.extend_from_slice(b"\\n"),
            0x0C => buf.extend_from_slice(b"\\f"),
            0x0D => buf.extend_from_slice(b"\\r"),
            b if b < 0x20 => {
                buf.extend_from_slice(b"\\u00");
                buf.push(hex_digit(b >> 4));
                buf.push(hex_digit(b & 0x0F));
            }
            _ => buf.push(b),
        }
    }
    buf.push(b'"');
}

fn hex_digit(b: u8) -> u8 {
    if b < 10 {
        b'0' + b
    } else {
        b'a' + (b - 10)
    }
}

/// Sort members by key using UTF-16 code-unit ordering (RFC 8785 §3.2.3).
fn serialize_object(buf: &mut Vec<u8>, members: &[Member]) -> Result<(), JcsError> {
    let mut sorted: Vec<&Member> = members.iter().collect();
    sorted.sort_by(|a, b| compare_utf16(&a.key, &b.key));

    buf.push(b'{');
    for (i, m) in sorted.iter().enumerate() {
        if i > 0 {
            buf.push(b',');
        }
        serialize_string(buf, &m.key);
        buf.push(b':');
        serialize_value(buf, &m.value)?;
    }
    buf.push(b'}');
    Ok(())
}

/// Compare two strings by their UTF-16 code unit sequences.
fn compare_utf16(a: &str, b: &str) -> std::cmp::Ordering {
    let mut a_iter = a.encode_utf16();
    let mut b_iter = b.encode_utf16();
    loop {
        match (a_iter.next(), b_iter.next()) {
            (None, None) => return std::cmp::Ordering::Equal,
            (None, Some(_)) => return std::cmp::Ordering::Less,
            (Some(_), None) => return std::cmp::Ordering::Greater,
            (Some(au), Some(bu)) => {
                if au != bu {
                    return au.cmp(&bu);
                }
            }
        }
    }
}
