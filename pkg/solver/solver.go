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

		if err == nil && len(aiGroups) == 4 {
			// Got all 4 groups from AI - perfect!
			return aiGroups, nil
		} else if err == nil && len(aiGroups) > 0 && len(aiGroups) < 4 {
			// Got partial results from AI - try to complete with pattern matching
			fmt.Printf("AI found %d of 4 groups. Trying pattern matching for remaining words...\n", len(aiGroups))

			// Find which words are already grouped
			usedWords := make(map[string]bool)
			for _, group := range aiGroups {
				for _, word := range group.Words {
					usedWords[word] = true
				}
			}

			// Get remaining words
			remainingWords := []string{}
			for _, word := range words {
				if !usedWords[word] {
					remainingWords = append(remainingWords, word)
				}
			}

			// Try pattern matching on remaining words
			if len(remainingWords) > 0 {
				patternGroups, _ := s.solveWithPatterns(remainingWords)

				// Combine AI groups with pattern groups
				allGroups := append(aiGroups, patternGroups...)

				if len(allGroups) == 4 {
					fmt.Printf("Successfully completed puzzle: %d AI groups + %d pattern groups\n", len(aiGroups), len(patternGroups))
					return allGroups, nil
				}
				// Return partial results
				return allGroups, fmt.Errorf("could only find %d groups", len(allGroups))
			}

			// Just return what AI found
			return aiGroups, fmt.Errorf("could only find %d groups", len(aiGroups))
		}
		// If AI fails completely, fall back to pattern matching
		if err != nil {
			fmt.Printf("AI analysis failed (%v), falling back to pattern matching...\n\n", err)
		}
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
	fmt.Printf("DEBUG solveWithPatterns: Called with %d words: %v\n", len(words), words)

	// Use the grouper to find potential groups
	candidates := s.grouper.FindGroups(words)
	fmt.Printf("DEBUG solveWithPatterns: Grouper found %d candidates\n", len(candidates))

	// Calculate how many groups we expect based on word count
	// Each group has 4 words, so expected groups = words / 4
	expectedGroups := len(words) / 4
	if expectedGroups > 4 {
		expectedGroups = 4 // Cap at 4 for standard Connections puzzle
	}
	fmt.Printf("DEBUG solveWithPatterns: Expecting %d groups from %d words\n", expectedGroups, len(words))

	// Find non-overlapping groups
	var result []Group
	used := make(map[string]bool)

	for i, candidate := range candidates {
		// Check if any word is already used
		hasUsed := false
		for _, word := range candidate.Words {
			if used[word] {
				hasUsed = true
				break
			}
		}

		if !hasUsed && len(result) < expectedGroups {
			fmt.Printf("DEBUG solveWithPatterns: Adding candidate %d: %v (theme: %s)\n", i, candidate.Words, candidate.Theme)
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
		} else if hasUsed {
			fmt.Printf("DEBUG solveWithPatterns: Skipping candidate %d (has used words)\n", i)
		}

		if len(result) == expectedGroups {
			break
		}
	}

	fmt.Printf("DEBUG solveWithPatterns: Found %d groups (expected %d)\n", len(result), expectedGroups)

	// Only return error if we're working with a full 16-word puzzle
	// For partial word sets (from AI completion), return what we found
	if len(words) == 16 && len(result) < 4 {
		return result, fmt.Errorf("could only find %d groups", len(result))
	}

	return result, nil
}
