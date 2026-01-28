package solver

import (
	"connections/pkg/ai"
	"connections/pkg/grouper"
	"fmt"
)

// Group represents a potential grouping of words
type Group struct {
	Words       []string
	Theme       string
	Explanation string
	Confidence  float64
	Source      string // "ai" or "pattern"
}

// Solver handles the logic for solving Connections puzzles
type Solver struct {
	grouper    *grouper.Grouper
	aiProvider ai.Provider
	useAI      bool
}

// New creates a new Solver instance with pattern matching only
func New() *Solver {
	return &Solver{
		grouper: grouper.New(),
		useAI:   false,
	}
}

// NewWithAI creates a new Solver instance with AI support
func NewWithAI(apiKey string) *Solver {
	return &Solver{
		grouper:    grouper.New(),
		aiProvider: ai.NewOpenAIProvider(apiKey),
		useAI:      true,
	}
}

// NewWithClaude creates a new Solver instance with Claude AI support
func NewWithClaude(apiKey string) *Solver {
	return &Solver{
		grouper:    grouper.New(),
		aiProvider: ai.NewClaudeProvider(apiKey),
		useAI:      true,
	}
}

// NewWithGemini creates a new Solver instance with Google Gemini AI support
func NewWithGemini(apiKey string) *Solver {
	return &Solver{
		grouper:    grouper.New(),
		aiProvider: ai.NewGeminiProvider(apiKey),
		useAI:      true,
	}
}

// Solve attempts to find the 4 groups from the 16 words
func (s *Solver) Solve(words []string) ([]Group, error) {
	if len(words) != 16 {
		return nil, fmt.Errorf("expected 16 words, got %d", len(words))
	}

	// Try AI first if enabled
	if s.useAI && s.aiProvider != nil {
		aiGroups, err := s.solveWithAI(words)
		if err == nil {
			return aiGroups, nil
		}
		// If AI fails, fall back to pattern matching
		fmt.Printf("AI analysis failed (%v), falling back to pattern matching...\n\n", err)
	}

	// Use pattern matching
	return s.solveWithPatterns(words)
}

// solveWithAI uses AI to find groups
func (s *Solver) solveWithAI(words []string) ([]Group, error) {
	suggestions, err := s.aiProvider.AnalyzeWords(words)
	if err != nil {
		return nil, err
	}

	var result []Group
	for _, suggestion := range suggestions {
		result = append(result, Group{
			Words:       suggestion.Words,
			Theme:       suggestion.Theme,
			Explanation: suggestion.Explanation,
			Confidence:  suggestion.Confidence,
			Source:      "ai",
		})
	}

	return result, nil
}

// solveWithPatterns uses pattern matching to find groups
func (s *Solver) solveWithPatterns(words []string) ([]Group, error) {
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
				Words:       candidate.Words,
				Theme:       candidate.Theme,
				Explanation: "",
				Confidence:  candidate.Confidence,
				Source:      "pattern",
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
