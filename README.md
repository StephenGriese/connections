# Connections Solver

A Go-based solver for the NYT Connections puzzle game with AI integration.

## Setup

### Prerequisites
- Go 1.21 or higher
- Anthropic API key (for Claude AI integration)

### Installation

1. Clone the repository:
```bash
git clone git@github.com-sjg:StephenGriese/connections.git
cd connections
```

2. Set up your Anthropic API key:
```bash
# Add to your ~/.zshrc or ~/.bashrc
export ANTHROPIC_API_KEY='your-api-key-here'

# Or source it for the current session
export ANTHROPIC_API_KEY='your-api-key-here'
```

3. Build and install:
```bash
make build
./build-and-install.sh
```

## Usage

Run the solver:
```bash
connections
```

Or from the project directory:
```bash
./run.sh
```

## Development

### Build
```bash
make build
```

### Test
```bash
make test
```

### Lint
```bash
make lint
```

## How It Works

The solver uses Claude AI to analyze the 16 words and identify the 4 groups of 4 words that share a common theme.

## License

MIT
