package analyzer

import (
	"testing"
)

func TestHasCommonPrefix(t *testing.T) {
	a := New()

	tests := []struct {
		name     string
		words    []string
		length   int
		expected string
		hasIt    bool
	}{
		{
			name:     "same prefix",
			words:    []string{"FIRE", "FIRST", "FIRM"},
			length:   2,
			expected: "fi",
			hasIt:    true,
		},
		{
			name:     "no common prefix",
			words:    []string{"FIRE", "WATER", "AIR"},
			length:   2,
			expected: "",
			hasIt:    false,
		},
		{
			name:     "empty list",
			words:    []string{},
			length:   2,
			expected: "",
			hasIt:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prefix, hasIt := a.HasCommonPrefix(tt.words, tt.length)
			if hasIt != tt.hasIt {
				t.Errorf("expected hasIt=%v, got %v", tt.hasIt, hasIt)
			}
			if prefix != tt.expected {
				t.Errorf("expected prefix=%q, got %q", tt.expected, prefix)
			}
		})
	}
}

func TestAllSameLength(t *testing.T) {
	a := New()

	tests := []struct {
		name     string
		words    []string
		expected bool
	}{
		{
			name:     "all same length",
			words:    []string{"FIRE", "WATER", "EARTH"},
			expected: false,
		},
		{
			name:     "all same length - true",
			words:    []string{"CAT", "DOG", "BAT"},
			expected: true,
		},
		{
			name:     "empty list",
			words:    []string{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := a.AllSameLength(tt.words)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
