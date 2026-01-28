# NYTimes Connections Solver

A Go-based solver for the NYTimes Connections puzzle.

## What is Connections?

Connections is a word puzzle where you're given 16 words that form 4 groups of 4 words each. Each group shares a common theme or connection. The challenge is to identify all four groups without making more than 4 mistakes.

## Project Structure

- `cmd/cli/` - Command-line interface
- `pkg/solver/` - Core solving logic
- `pkg/grouper/` - Group detection algorithms
- `pkg/analyzer/` - Word analysis utilities

## Usage

### CLI
```bash
go run ./cmd/cli/main.go
```

## Development

Build the project:
```bash
make build
```

Run tests:
```bash
make test
```

## Future Plans

- Web interface
- AI-assisted solving using LLM APIs
- Pattern recognition improvements
- Learning from past puzzles
