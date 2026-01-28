#!/bin/bash

# Test the Connections solver with Claude AI

# Make sure ANTHROPIC_API_KEY is set in your environment
if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo "Error: ANTHROPIC_API_KEY environment variable not set"
    echo "Set it in your ~/.zshrc or run: export ANTHROPIC_API_KEY='your-key-here'"
    exit 1
fi

echo "🔗 Testing Connections Solver with Claude AI"
echo "============================================"
echo "API Key set: ${ANTHROPIC_API_KEY:0:20}..."
echo ""

echo "Building the project..."
make build

echo ""
echo "Testing with example puzzle..."
echo ""
echo "BASS CLUB DIAMOND HEART SPADE SOLE PERCH TROUT WOOD IRON DRIVER PUTTER ACE KING QUEEN JACK" | ./target/local/bin/connections
