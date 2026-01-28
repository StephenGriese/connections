package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"connections/pkg/solver"
)

func main() {
	fmt.Println("🔗 NYTimes Connections Solver")
	fmt.Println("================================")
	fmt.Println()

	words, err := readWords()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading words: %v\n", err)
		os.Exit(1)
	}

	if len(words) != 16 {
		fmt.Fprintf(os.Stderr, "Error: Expected 16 words, got %d\n", len(words))
		os.Exit(1)
	}

	fmt.Println("Words entered:")
	for i, word := range words {
		fmt.Printf("%2d. %s\n", i+1, word)
	}
	fmt.Println()

	// Solve the puzzle
	s := solver.New()
	groups, err := s.Solve(words)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error solving: %v\n", err)
		os.Exit(1)
	}

	// Display results
	fmt.Println("Suggested Groups:")
	fmt.Println("=================")
	for i, group := range groups {
		fmt.Printf("\nGroup %d: %s\n", i+1, group.Theme)
		fmt.Printf("Words: %s\n", strings.Join(group.Words, ", "))
		fmt.Printf("Confidence: %.0f%%\n", group.Confidence*100)
	}
}

func readWords() ([]string, error) {
	fmt.Println("Enter 16 words (one per line, or all on one line separated by spaces/commas):")

	scanner := bufio.NewScanner(os.Stdin)
	var words []string

	// Read first line
	if !scanner.Scan() {
		return nil, scanner.Err()
	}

	line := scanner.Text()

	// Check if all words are on one line
	line = strings.ReplaceAll(line, ",", " ")
	parts := strings.Fields(line)

	if len(parts) == 16 {
		return parts, nil
	} else if len(parts) == 1 {
		words = append(words, parts[0])
	} else if len(parts) > 1 {
		words = append(words, parts...)
	}

	// Read remaining lines until we have 16 words
	for len(words) < 16 && scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", " ")
		parts := strings.Fields(line)
		words = append(words, parts...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
