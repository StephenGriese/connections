package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"connections/pkg/solver"
)

func main() {
	// Try to load .env file if it exists (ignore errors if not found)
	loadEnvFile()

	fmt.Println("🔗 NYTimes Connections Solver")
	fmt.Println("================================")

	// Check for AI API keys (Gemini first, then Claude, then OpenAI)
	geminiKey := os.Getenv("GEMINI_API_KEY")
	claudeKey := os.Getenv("ANTHROPIC_API_KEY")
	openaiKey := os.Getenv("OPENAI_API_KEY")

	var aiMode string
	if geminiKey != "" {
		aiMode = "gemini"
		fmt.Println("✨ AI mode enabled (using Google Gemini)")
	} else if claudeKey != "" {
		aiMode = "claude"
		fmt.Println("✨ AI mode enabled (using Claude)")
	} else if openaiKey != "" {
		aiMode = "openai"
		fmt.Println("✨ AI mode enabled (using OpenAI)")
	} else {
		aiMode = "none"
		fmt.Println("📊 Pattern matching mode")
		fmt.Println("   Set GEMINI_API_KEY, ANTHROPIC_API_KEY, or OPENAI_API_KEY for AI")
	}
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

	// Create solver (with or without AI)
	var s *solver.Solver
	switch aiMode {
	case "gemini":
		s = solver.NewWithGemini(geminiKey)
	case "claude":
		s = solver.NewWithClaude(claudeKey)
	case "openai":
		s = solver.NewWithAI(openaiKey)
	default:
		s = solver.New()
	}

	// Solve the puzzle
	groups, err := s.Solve(words)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error solving: %v\n", err)
		os.Exit(1)
	}

	// Display results
	fmt.Println("Suggested Groups:")
	fmt.Println("=================")
	for i, group := range groups {
		fmt.Printf("\nGroup %d: %s", i+1, group.Theme)
		if group.Source == "ai" {
			fmt.Printf(" [AI]")
		} else {
			fmt.Printf(" [Pattern]")
		}
		fmt.Println()
		fmt.Printf("Words: %s\n", strings.Join(group.Words, ", "))
		if group.Explanation != "" {
			fmt.Printf("Explanation: %s\n", group.Explanation)
		}
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

// loadEnvFile loads environment variables from .env file if it exists
func loadEnvFile() {
	file, err := os.Open(".env")
	if err != nil {
		return // File doesn't exist, that's okay
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Remove quotes if present
			value = strings.Trim(value, "\"'")

			// Only set if not already set in environment
			if os.Getenv(key) == "" {
				_ = os.Setenv(key, value)
			}
		}
	}
}
