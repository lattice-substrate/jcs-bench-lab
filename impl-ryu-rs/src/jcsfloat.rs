use crate::jcserr::{FailureClass, JcsError};
use crate::jcsfloat_format::format_ecma;
use crate::pow10_table::{POW10_MAX_EXP, POW10_MIN_EXP, POW10_TABLE};

const MANT_BITS_64: u32 = 52;
const EXP_BITS_64: u32 = 11;
const BIAS_64: i32 = 1023;
const MANT_MASK_64: u64 = (1u64 << MANT_BITS_64) - 1;
const EXP_MASK_64: u32 = (1u32 << EXP_BITS_64) - 1;

/// Formats an IEEE 754 double-precision value exactly as specified by the
/// ECMAScript Number::toString algorithm (ECMA-262, radix 10).
///
/// Uses the Ryu algorithm (Ulf Adams, PLDI 2018) with fixed-width 128-bit
/// arithmetic and lastRemovedDigit-based shortest-output extraction.
pub fn format_double(f: f64) -> Result<String, JcsError> {
    // ECMA-FMT-001: NaN -> error
    if f.is_nan() {
        return Err(JcsError::new(
            FailureClass::InvalidGrammar,
            -1,
            "NaN is not representable in JSON",
        ));
    }
    // ECMA-FMT-002: -0 and +0 -> "0"
    if f == 0.0 {
        return Ok("0".to_string());
    }
    // ECMA-FMT-003: +/-Infinity -> error
    if f.is_infinite() {
        return Err(JcsError::new(
            FailureClass::InvalidGrammar,
            -1,
            "Infinity is not representable in JSON",
        ));
    }

    let negative;
    let fabs;
    if f < 0.0 {
        negative = true;
        fabs = -f;
    } else {
        negative = false;
        fabs = f;
    }

    let (digits, dp) = ryu_float64(fabs);
    Ok(format_ecma(negative, &digits, dp))
}

fn ryu_float64(f: f64) -> (String, i32) {
    let fbits = f.to_bits();
    let raw_mant = fbits & MANT_MASK_64;
    let raw_exp = ((fbits >> MANT_BITS_64) as u32) & EXP_MASK_64;

    let mant: u64;
    let exp: i32;

    if raw_exp == 0 {
        mant = raw_mant;
        exp = 1 - BIAS_64 - MANT_BITS_64 as i32;
    } else {
        mant = (1u64 << MANT_BITS_64) | raw_mant;
        exp = raw_exp as i32 - BIAS_64 - MANT_BITS_64 as i32;
    }

    // Fast path: exact integer.
    if exp <= 0 && mant.trailing_zeros() >= (-exp) as u32 {
        let m = mant >> ((-exp) as u32);
        return (format_uint(m), count_decimal_digits(m));
    }

    // Ryu interval computation.
    let (ml, mc, mu, e2) = compute_interval(mant, exp);
    if e2 == 0 {
        return ryu_shortest(ml, mc, mu, true, false, 0);
    }

    let q = mul_by_log2_log10(-e2) + 1;

    let (dl, _, dl0_init) = mult_128bit_pow10(ml, e2, q);
    let (dc, _, dc0_init) = mult_128bit_pow10(mc, e2, q);
    let (du, e2_final, du0_init) = mult_128bit_pow10(mu, e2, q);

    if e2_final >= 0 {
        return ("0".to_string(), 0);
    }

    let mut dl0 = dl0_init;
    let mut dc0 = dc0_init;
    let mut du0 = du0_init;

    if q > 55 {
        dl0 = false;
        dc0 = false;
        du0 = false;
    }
    if (-24..0).contains(&q) {
        if divisible_by_power5(ml, (-q) as u32) {
            dl0 = true;
        }
        if divisible_by_power5(mc, (-q) as u32) {
            dc0 = true;
        }
        if divisible_by_power5(mu, (-q) as u32) {
            du0 = true;
        }
    }

    let extra = (-e2_final) as u32;
    let extra_mask = (1u64 << extra) - 1;

    let fracl = dl & extra_mask;
    let fracc = dc & extra_mask;
    let fracu = du & extra_mask;
    let mut dl = dl >> extra;
    let dc = dc >> extra;
    let mut du = du >> extra;

    let accept_bounds = mant & 1 == 0;
    let mut uok = !du0 || fracu > 0;
    if du0 && fracu == 0 {
        uok = accept_bounds;
    }
    if !uok {
        du -= 1;
    }

    let cup = if dc0 {
        fracc > (1u64 << (extra - 1))
            || (fracc == (1u64 << (extra - 1)) && dc & 1 == 1)
    } else {
        (fracc >> (extra - 1)) == 1
    };

    let lok = dl0 && fracl == 0 && accept_bounds;
    if !lok {
        dl += 1;
    }

    let c0 = dc0 && fracc == 0;

    ryu_shortest(dl, dc, du, c0, cup, q)
}

fn compute_interval(mant: u64, exp: i32) -> (u64, u64, u64, i32) {
    if mant != (1u64 << MANT_BITS_64) || exp == 1 - BIAS_64 - MANT_BITS_64 as i32 {
        (2 * mant - 1, 2 * mant, 2 * mant + 1, exp - 1)
    } else {
        (4 * mant - 1, 4 * mant, 4 * mant + 2, exp - 2)
    }
}

fn ryu_shortest(
    mut dl: u64,
    mut dc: u64,
    mut du: u64,
    mut c0: bool,
    mut cup: bool,
    q: i32,
) -> (String, i32) {
    let mut trimmed: i32 = 0;
    let mut c_next_digit: u64 = 0;

    while du > 0 {
        let l = dl.div_ceil(10);
        let c = dc / 10;
        let cdigit = dc % 10;
        let u = du / 10;
        if l > u {
            break;
        }
        if l == c + 1 && c < u {
            dc = c + 1;
            dl = l;
            du = u;
            c_next_digit = 0;
            cup = false;
            trimmed += 1;
            c0 = c0 && c_next_digit == 0;
            continue;
        }
        trimmed += 1;
        c0 = c0 && c_next_digit == 0;
        c_next_digit = cdigit;
        dl = l;
        dc = c;
        du = u;
    }

    if trimmed > 0 {
        cup = c_next_digit > 5
            || c_next_digit == 5 && !c0
            || c_next_digit == 5 && dc & 1 == 1;
    }
    if dc < du && cup {
        dc += 1;
    }

    let mut buf = [0u8; 20];
    let mut n = 20usize;
    let mut v = dc;
    while v > 0 {
        n -= 1;
        buf[n] = b'0' + (v % 10) as u8;
        v /= 10;
    }
    if n == 20 {
        n = 19;
        buf[19] = b'0';
    }

    let ndigits = (20 - n) as i32;

    let mut end = 20usize;
    while end > n + 1 && buf[end - 1] == b'0' {
        end -= 1;
        trimmed += 1;
    }

    let dp = ndigits + trimmed - q;
    let digits = std::str::from_utf8(&buf[n..end]).unwrap().to_string();
    (digits, dp)
}

fn mul_by_log2_log10(x: i32) -> i32 {
    ((x as i64 * 78913) >> 18) as i32
}

fn mul_by_log10_log2(x: i32) -> i32 {
    ((x as i64 * 108853) >> 15) as i32
}

fn mult_128bit_pow10(m: u64, e2: i32, q: i32) -> (u64, i32, bool) {
    if q == 0 {
        return (m << 8, e2 - 8, true);
    }
    if !(POW10_MIN_EXP..=POW10_MAX_EXP).contains(&q) {
        return (0, 0, false);
    }
    let idx = (q - POW10_MIN_EXP) as usize;
    let mut pow = POW10_TABLE[idx];
    if q < 0 {
        pow.0 = pow.0.wrapping_add(1);
    }
    let e2_out = e2 + mul_by_log10_log2(q) - 127 + 119;

    let full_lo = m as u128 * pow.0 as u128;
    let full_hi = m as u128 * pow.1 as u128;
    let l1 = (full_lo >> 64) as u64;
    let l0 = full_lo as u64;
    let h1 = (full_hi >> 64) as u64;
    let h0 = full_hi as u64;
    let (mid, carry) = l1.overflowing_add(h0);
    let h1 = h1 + carry as u64;

    (
        (h1 << 9) | (mid >> 55),
        e2_out,
        (mid << 9) == 0 && l0 == 0,
    )
}

fn divisible_by_power5(mut m: u64, k: u32) -> bool {
    if m == 0 {
        return true;
    }
    for _ in 0..k {
        if m % 5 != 0 {
            return false;
        }
        m /= 5;
    }
    true
}

fn format_uint(mut v: u64) -> String {
    if v == 0 {
        return "0".to_string();
    }
    let mut buf = [0u8; 20];
    let mut n = 20usize;
    while v > 0 {
        n -= 1;
        buf[n] = b'0' + (v % 10) as u8;
        v /= 10;
    }
    let mut end = 20usize;
    while end > n + 1 && buf[end - 1] == b'0' {
        end -= 1;
    }
    std::str::from_utf8(&buf[n..end]).unwrap().to_string()
}

fn count_decimal_digits(mut v: u64) -> i32 {
    if v == 0 {
        return 1;
    }
    let mut n = 0i32;
    while v > 0 {
        n += 1;
        v /= 10;
    }
    n
}
