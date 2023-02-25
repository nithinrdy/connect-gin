package utilities

import "unicode"

func IsPasswordValid(s string) bool {
	var (
		rightLen  = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(s) >= 7 && len(s) <= 15 {
		rightLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true

		}
	}
	return rightLen && hasUpper && hasLower && hasNumber
}
