# Quick Start Guide

## Installation

```bash
cd ~/repos/sjg/connections
make build
```

This will create the executable at `target/local/bin/connections`.

## Running the Solver

### Using the built binary:
```bash
./target/local/bin/connections
```

### Using `go run`:
```bash
make run
# or
go run ./cmd/cli/main.go
```

## Usage

When you run the program, it will prompt you to enter 16 words. You can either:

1. **Enter one word per line** (press Enter after each word)
2. **Enter all words on one line** separated by spaces or commas

### Example Session:

```
🔗 NYTimes Connections Solver
================================

Enter 16 words (one per line, or all on one line separated by spaces/commas):
FIRE FIRST FIRM FISH BLUE BLUR BLURT BLUSH APPLE APPLY APPS APRON DOG CAT BIRD RAT

Words entered:
 1. FIRE
 2. FIRST
 3. FIRM
 4. FISH
 5. BLUE
 6. BLUR
 7. BLURT
 8. BLUSH
 9. APPLE
10. APPLY
11. APPS
12. APRON
13. DOG
14. CAT
15. BIRD
16. RAT

Suggested Groups:
=================

Group 1: Words starting with 'fir'
Words: FIRE, FIRST, FIRM
Confidence: 50%

Group 2: Words starting with 'blu'
Words: BLUE, BLUR, BLURT, BLUSH
Confidence: 50%

...
```

## Current Limitations

The current implementation uses basic pattern matching:
- Common prefixes
- Common suffixes
- Length patterns
- Basic compound word detection

It may not find all 4 groups for complex puzzles that require:
- Semantic/categorical knowledge
- Cultural references
- Wordplay and puns

## Future Improvements

See `ARCHITECTURE.md` for planned enhancements, including AI-assisted solving.
