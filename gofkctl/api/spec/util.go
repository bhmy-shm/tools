package spec

import "strings"

func Untitle(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]
}

// Title returns a string value with s[0] which has been convert into upper case that
// there are not empty input text
func Title(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

// Contains checks if str is in list.
func Contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}

	return false
}
