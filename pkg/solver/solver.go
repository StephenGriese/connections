package solver

import (
	"connections/pkg/grouper"
	"fmt"
)

// Group represents a potential grouping of words
type Group struct {
	Words      []string
	Theme      string
	Confidence float64
}

// Solver handles the logic for solving Connections puzzles
type Solver struct {
	grouper *grouper.Grouper
}

// New creates a new Solver instance
func New() *Solver {
	return &Solver{
		grouper: grouper.New(),
	}
}

// Solve attempts to find the 4 groups from the 16 words
func (s *Solver) Solve(words []string) ([]Group, error) {
	if len(words) != 16 {
		return nil, fmt.Errorf("expected 16 words, got %d", len(words))
	}

	// Use the grouper to find potential groups
	candidates := s.grouper.FindGroups(words)

	// For now, return the top 4 candidates
	// In a more sophisticated version, we'd use constraint satisfaction
	// to ensure exactly 4 non-overlapping groups
	var result []Group
	used := make(map[string]bool)

	for _, candidate := range candidates {
		// Check if any word is already used
		hasUsed := false
		for _, word := range candidate.Words {
			if used[word] {
				hasUsed = true
				break
			}
		}

		if !hasUsed && len(result) < 4 {
			result = append(result, Group{
				Words:      candidate.Words,
				Theme:      candidate.Theme,
				Confidence: candidate.Confidence,
			})

			// Mark words as used
			for _, word := range candidate.Words {
				used[word] = true
			}
		}

		if len(result) == 4 {
			break
		}
	}

	if len(result) < 4 {
		return result, fmt.Errorf("could only find %d groups", len(result))
	}

	return result, nil
}
