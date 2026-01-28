package analyzer

import (
	"strings"
)

// Analyzer provides word analysis utilities
type Analyzer struct {
}

// New creates a new Analyzer instance
func New() *Analyzer {
	return &Analyzer{}
}

// HasCommonPrefix checks if words share a common prefix of given length
func (a *Analyzer) HasCommonPrefix(words []string, length int) (string, bool) {
	if len(words) == 0 || length == 0 {
		return "", false
	}

	prefix := ""
	if len(words[0]) >= length {
		prefix = strings.ToLower(words[0][:length])
	} else {
		return "", false
	}

	for _, word := range words[1:] {
		if len(word) < length {
			return "", false
		}
		if strings.ToLower(word[:length]) != prefix {
			return "", false
		}
	}

	return prefix, true
}

// HasCommonSuffix checks if words share a common suffix of given length
func (a *Analyzer) HasCommonSuffix(words []string, length int) (string, bool) {
	if len(words) == 0 || length == 0 {
		return "", false
	}

	suffix := ""
	if len(words[0]) >= length {
		suffix = strings.ToLower(words[0][len(words[0])-length:])
	} else {
		return "", false
	}

	for _, word := range words[1:] {
		if len(word) < length {
			return "", false
		}
		if strings.ToLower(word[len(word)-length:]) != suffix {
			return "", false
		}
	}

	return suffix, true
}

// AllSameLength checks if all words have the same length
func (a *Analyzer) AllSameLength(words []string) bool {
	if len(words) == 0 {
		return true
	}

	length := len(words[0])
	for _, word := range words[1:] {
		if len(word) != length {
			return false
		}
	}

	return true
}

// ContainsSubstring checks if a word contains a substring (case-insensitive)
func (a *Analyzer) ContainsSubstring(word, substring string) bool {
	return strings.Contains(strings.ToLower(word), strings.ToLower(substring))
}
