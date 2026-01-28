package grouper

import (
	"connections/pkg/analyzer"
	"sort"
	"strings"
)

// Candidate represents a potential group of 4 words
type Candidate struct {
	Words      []string
	Theme      string
	Confidence float64
}

// Grouper finds potential groupings of words
type Grouper struct {
	analyzer *analyzer.Analyzer
}

// New creates a new Grouper instance
func New() *Grouper {
	return &Grouper{
		analyzer: analyzer.New(),
	}
}

// FindGroups analyzes words and returns potential groupings sorted by confidence
func (g *Grouper) FindGroups(words []string) []Candidate {
	var candidates []Candidate

	// Strategy 1: Find words with common prefixes
	candidates = append(candidates, g.findPrefixGroups(words)...)

	// Strategy 2: Find words with common suffixes
	candidates = append(candidates, g.findSuffixGroups(words)...)

	// Strategy 3: Find words by length patterns
	candidates = append(candidates, g.findLengthGroups(words)...)

	// Strategy 4: Find compound word parts
	candidates = append(candidates, g.findCompoundGroups(words)...)

	// Sort by confidence (highest first)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Confidence > candidates[j].Confidence
	})

	return candidates
}

func (g *Grouper) findPrefixGroups(words []string) []Candidate {
	prefixMap := make(map[string][]string)

	for _, word := range words {
		if len(word) >= 3 {
			prefix := strings.ToLower(word[:3])
			prefixMap[prefix] = append(prefixMap[prefix], word)
		}
	}

	var candidates []Candidate
	for prefix, group := range prefixMap {
		if len(group) >= 4 {
			candidates = append(candidates, Candidate{
				Words:      group[:4],
				Theme:      "Words starting with '" + prefix + "'",
				Confidence: 0.5,
			})
		}
	}

	return candidates
}

func (g *Grouper) findSuffixGroups(words []string) []Candidate {
	suffixMap := make(map[string][]string)

	for _, word := range words {
		if len(word) >= 3 {
			suffix := strings.ToLower(word[len(word)-3:])
			suffixMap[suffix] = append(suffixMap[suffix], word)
		}
	}

	var candidates []Candidate
	for suffix, group := range suffixMap {
		if len(group) >= 4 {
			candidates = append(candidates, Candidate{
				Words:      group[:4],
				Theme:      "Words ending with '" + suffix + "'",
				Confidence: 0.5,
			})
		}
	}

	return candidates
}

func (g *Grouper) findLengthGroups(words []string) []Candidate {
	lengthMap := make(map[int][]string)

	for _, word := range words {
		lengthMap[len(word)] = append(lengthMap[len(word)], word)
	}

	var candidates []Candidate
	for _, group := range lengthMap {
		if len(group) == 4 {
			candidates = append(candidates, Candidate{
				Words:      group,
				Theme:      "All words have same length",
				Confidence: 0.3,
			})
		}
	}

	return candidates
}

func (g *Grouper) findCompoundGroups(words []string) []Candidate {
	// Look for words that could be parts of compound words
	// e.g., BLUE, BIRD could both go with "BERRY" -> BLUEBERRY, BLACKBIRD

	// This is a placeholder for more sophisticated compound word detection
	var candidates []Candidate

	// Common compound word patterns
	commonParts := []string{"FIRE", "WATER", "SNOW", "SUPER", "OVER", "UNDER"}

	for _, part := range commonParts {
		var matching []string
		for _, word := range words {
			upper := strings.ToUpper(word)
			if strings.Contains(upper, part) {
				matching = append(matching, word)
			}
		}

		if len(matching) >= 4 {
			candidates = append(candidates, Candidate{
				Words:      matching[:4],
				Theme:      "Related to '" + part + "'",
				Confidence: 0.4,
			})
		}
	}

	return candidates
}
