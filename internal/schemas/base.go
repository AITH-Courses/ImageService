package schemas

import (
	"unicode"
	"unicode/utf8"
)

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	if unicode.IsLower(r) {
		return string(unicode.ToTitle(r)) + s[size:]
	} else {
		return s
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: capitalize(message),
	}
}
