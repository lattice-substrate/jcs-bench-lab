/// Applies the ECMA-262 §6.1.6.1.20 formatting rules (steps 7-10).
///
/// `digits`: significand digit string (k digits)
/// `n`: decimal exponent (number of integer digits in fixed-point view)
pub fn format_ecma(negative: bool, digits: &str, n: i32) -> String {
    let k = digits.len() as i32;
    let mut buf = String::with_capacity(digits.len() + 8);

    if negative {
        buf.push('-');
    }

    if k <= n && n <= 21 {
        // ECMA-FMT-004: integer with trailing zeros
        buf.push_str(digits);
        for _ in 0..(n - k) {
            buf.push('0');
        }
    } else if 0 < n && n <= 21 {
        // ECMA-FMT-005: fixed decimal
        buf.push_str(&digits[..n as usize]);
        buf.push('.');
        buf.push_str(&digits[n as usize..]);
    } else if -6 < n && n <= 0 {
        // ECMA-FMT-006: 0.000...digits
        buf.push_str("0.");
        for _ in 0..(-n) {
            buf.push('0');
        }
        buf.push_str(digits);
    } else {
        // ECMA-FMT-007: exponential notation
        buf.push(digits.as_bytes()[0] as char);
        if k > 1 {
            buf.push('.');
            buf.push_str(&digits[1..]);
        }
        buf.push('e');
        let exp = n - 1;
        if exp >= 0 {
            buf.push('+');
        }
        append_int(&mut buf, exp);
    }

    buf
}

fn append_int(buf: &mut String, mut v: i32) {
    if v < 0 {
        buf.push('-');
        v = -v;
    }
    if v == 0 {
        buf.push('0');
        return;
    }
    let mut tmp = [0u8; 11];
    let mut i = tmp.len();
    while v > 0 {
        i -= 1;
        tmp[i] = b'0' + (v % 10) as u8;
        v /= 10;
    }
    buf.push_str(std::str::from_utf8(&tmp[i..]).unwrap());
}
