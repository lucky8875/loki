package strings

import "unicode"

// IsEmpty check if string of len is 0 or ""
func IsEmpty(s string) bool {
	return len(s) == 0 || s == ""
}

// IsNotEmpty check if string of len is not 0 or ""
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsBlank checks if string is empty ("") or whitespace only.
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	for _, v := range s {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}

// IsNotBlank checks if string is not empty ("") or whitespace only.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsPtrBlank checks if string pointer is nil or empty ("") or whitespace only.
func IsPtrBlank(s *string) bool {
	return s == nil || IsBlank(*s)
}

// IsPtrEmpty checks if string pointer is nil or string of len is 0 or ""
func IsPtrEmpty(s *string) bool {
	return s == nil || IsEmpty(*s)
}

// IsAlpha checks if string only contains unicode letters
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, v := range s {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {
	if s == "" {
		return false
	}
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}
