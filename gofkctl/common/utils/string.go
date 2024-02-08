package utils

import "strings"

var WhiteSpace = []rune{'\n', '\t', '\f', '\v', ' '}

func ContainsAny(s string, runes ...rune) bool {
	if len(runes) == 0 {
		return true
	}
	tmp := make(map[rune]struct{}, len(runes))
	for _, r := range runes {
		tmp[r] = struct{}{}
	}

	for _, r := range s {
		if _, ok := tmp[r]; ok {
			return true
		}
	}
	return false
}

func ContainsWhiteSpace(s string) bool {
	return ContainsAny(s, WhiteSpace...)
}

func TrimWhiteSpace(s string) string {
	r := strings.NewReplacer(" ", "", "\t", "", "\n", "", "\f", "", "\r", "")
	return r.Replace(s)
}

func IsEmptyStringOrWhiteSpace(s string) bool {
	v := TrimWhiteSpace(s)
	return len(v) == 0
}
