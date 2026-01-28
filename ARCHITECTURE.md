# NYTimes Connections Solver

## Architecture

This project is structured to be extensible and testable, following Go best practices.

### Directory Structure

```
connections/
├── cmd/cli/              # Command-line interface entry point
│   └── main.go
├── pkg/
│   ├── solver/          # Main solving logic
│   │   ├── solver.go
│   │   └── solver_test.go
│   ├── grouper/         # Group detection strategies
│   │   └── grouper.go
│   └── analyzer/        # Word analysis utilities
│       ├── analyzer.go
│       └── analyzer_test.go
├── go.mod
├── Makefile
└── README.md
```

### Components

#### Solver (`pkg/solver`)
The main orchestrator that coordinates the solving process. It takes 16 words and returns 4 groups of 4 words each, along with theme descriptions and confidence scores.

#### Grouper (`pkg/grouper`)
Implements various strategies for finding potential word groupings:
- Prefix matching (e.g., words starting with the same letters)
- Suffix matching (e.g., words ending with the same letters)
- Length patterns (e.g., all 4-letter words)
- Compound word detection (e.g., words that can combine with a common word)

#### Analyzer (`pkg/analyzer`)
Provides low-level word analysis utilities used by the Grouper:
- Common prefix/suffix detection
- Length comparison
- Substring matching

### Design Principles

1. **Separation of Concerns**: Each package has a single, well-defined responsibility
2. **Testability**: All core logic is in packages that can be easily unit tested
3. **Extensibility**: New grouping strategies can be added to the Grouper without changing other code
4. **Confidence Scoring**: Each suggested group has a confidence score to help prioritize

### Future Enhancements

- **AI Integration**: Add LLM-based analysis for semantic groupings
- **Category Detection**: Build knowledge base of common connection types
- **Learning System**: Improve from feedback on actual puzzle solutions
- **Web Interface**: Add a web UI similar to the Wordle Helper
- **Historical Data**: Learn from past NYTimes Connections puzzles
