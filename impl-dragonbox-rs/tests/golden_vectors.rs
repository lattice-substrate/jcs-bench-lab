use jcs_dragonbox_rs::jcsfloat::format_double;
use std::fs;
use std::path::PathBuf;

fn testdata_dir() -> PathBuf {
    PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("testdata")
}

fn run_vector_file(filename: &str) {
    let path = testdata_dir().join(filename);
    let data = fs::read_to_string(&path)
        .unwrap_or_else(|e| panic!("failed to read {}: {}", path.display(), e));

    let mut pass = 0usize;
    let mut fail = 0usize;
    let mut first_failure: Option<String> = None;

    for (lineno, line) in data.lines().enumerate() {
        let line = line.trim();
        if line.is_empty() {
            continue;
        }
        let (hex_str, expected) = match line.split_once(',') {
            Some(pair) => pair,
            None => panic!("{}:{}: malformed line: {}", filename, lineno + 1, line),
        };

        let bits = u64::from_str_radix(hex_str, 16)
            .unwrap_or_else(|e| panic!("{}:{}: bad hex '{}': {}", filename, lineno + 1, hex_str, e));
        let f = f64::from_bits(bits);

        match format_double(f) {
            Ok(got) => {
                if got == expected {
                    pass += 1;
                } else {
                    fail += 1;
                    if first_failure.is_none() {
                        first_failure = Some(format!(
                            "{}:{}: bits=0x{:016x} expected={} got={}",
                            filename,
                            lineno + 1,
                            bits,
                            expected,
                            got
                        ));
                    }
                }
            }
            Err(e) => {
                // Zero is special — bits 0x0000000000000000 should produce "0"
                fail += 1;
                if first_failure.is_none() {
                    first_failure = Some(format!(
                        "{}:{}: bits=0x{:016x} expected={} error={}",
                        filename,
                        lineno + 1,
                        bits,
                        expected,
                        e
                    ));
                }
            }
        }
    }

    if let Some(msg) = first_failure {
        panic!(
            "{}: {}/{} vectors failed. First failure: {}",
            filename,
            fail,
            pass + fail,
            msg
        );
    }
    eprintln!("{}: {}/{} vectors passed", filename, pass, pass);
}

#[test]
fn golden_vectors() {
    run_vector_file("golden_vectors.csv");
}

#[test]
fn golden_stress_vectors() {
    run_vector_file("golden_stress_vectors.csv");
}
