#!/bin/bash

# Run the Connections solver with Claude AI

# Make sure ANTHROPIC_API_KEY is set in your environment
# Example: export ANTHROPIC_API_KEY='your-key-here'

if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo "Warning: ANTHROPIC_API_KEY environment variable not set"
    echo "Set it in your ~/.zshrc or run: export ANTHROPIC_API_KEY='your-key-here'"
fi

# Run the solver
./target/local/bin/connections
