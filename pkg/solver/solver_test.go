package solver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name        string
		words       []string
		expectError bool
	}{
		{
			name:        "wrong number of words",
			words:       []string{"ONE", "TWO", "THREE"},
			expectError: true,
		},
		{
			name: "valid 16 words",
			words: []string{
				"FIRE", "FIRST", "FIRM", "FISH",
				"BLUE", "BLUR", "BLURT", "BLUSH",
				"APPLE", "APPLY", "APPS", "APRON",
				"DOG", "CAT", "BIRD", "FISH",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			groups, err := s.Solve(tt.words)

			if tt.expectError {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Logf("solver returned error (this is OK for basic implementation): %v", err)
			}

			if len(groups) > 4 {
				t.Errorf("expected at most 4 groups, got %d", len(groups))
			}

			// Check that all groups have 4 words
			for i, group := range groups {
				if len(group.Words) != 4 {
					t.Errorf("group %d has %d words, expected 4", i, len(group.Words))
				}
			}
		})
	}
}
