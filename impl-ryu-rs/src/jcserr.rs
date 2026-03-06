use std::fmt;

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum FailureClass {
    InvalidUtf8,
    InvalidGrammar,
    DuplicateKey,
    LoneSurrogate,
    Noncharacter,
    NumberOverflow,
    NumberNegZero,
    NumberUnderflow,
    BoundExceeded,
    NotCanonical,
    CliUsage,
    InternalIo,
    InternalError,
}

impl FailureClass {
    pub fn exit_code(&self) -> i32 {
        match self {
            FailureClass::InternalIo | FailureClass::InternalError => 10,
            _ => 2,
        }
    }

    pub fn as_str(&self) -> &'static str {
        match self {
            FailureClass::InvalidUtf8 => "INVALID_UTF8",
            FailureClass::InvalidGrammar => "INVALID_GRAMMAR",
            FailureClass::DuplicateKey => "DUPLICATE_KEY",
            FailureClass::LoneSurrogate => "LONE_SURROGATE",
            FailureClass::Noncharacter => "NONCHARACTER",
            FailureClass::NumberOverflow => "NUMBER_OVERFLOW",
            FailureClass::NumberNegZero => "NUMBER_NEGZERO",
            FailureClass::NumberUnderflow => "NUMBER_UNDERFLOW",
            FailureClass::BoundExceeded => "BOUND_EXCEEDED",
            FailureClass::NotCanonical => "NOT_CANONICAL",
            FailureClass::CliUsage => "CLI_USAGE",
            FailureClass::InternalIo => "INTERNAL_IO",
            FailureClass::InternalError => "INTERNAL_ERROR",
        }
    }
}

#[derive(Debug, Clone)]
pub struct JcsError {
    pub class: FailureClass,
    pub offset: i64,
    pub message: String,
}

impl JcsError {
    pub fn new(class: FailureClass, offset: i64, message: impl Into<String>) -> Self {
        Self {
            class,
            offset,
            message: message.into(),
        }
    }

    pub fn exit_code(&self) -> i32 {
        self.class.exit_code()
    }
}

impl fmt::Display for JcsError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        if self.offset >= 0 {
            write!(
                f,
                "jcserr: {} at byte {}: {}",
                self.class.as_str(),
                self.offset,
                self.message
            )
        } else {
            write!(f, "jcserr: {}: {}", self.class.as_str(), self.message)
        }
    }
}

impl std::error::Error for JcsError {}
