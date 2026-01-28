
#!/bin/bash

# Demo script showing both pattern-matching and AI modes

echo "========================================="
echo "Connections Solver Demo"
echo "========================================="
echo ""

# Example puzzle (from EXAMPLES.md)
PUZZLE="BASS CLUB DIAMOND HEART SPADE SOLE PERCH TROUT WOOD IRON DRIVER PUTTER ACE KING QUEEN JACK"

echo "Example puzzle:"
echo "$PUZZLE"
echo ""
echo "Expected groups:"
echo "  - Card suits: CLUB, DIAMOND, HEART, SPADE"
echo "  - Fish: BASS, SOLE, PERCH, TROUT"
echo "  - Golf clubs: WOOD, IRON, DRIVER, PUTTER"
echo "  - Face cards: ACE, KING, QUEEN, JACK"
echo ""
echo "========================================="
echo ""

# Check if OpenAI API key is set
if [ -z "$OPENAI_API_KEY" ]; then
    echo "⚠️  OPENAI_API_KEY not set - using pattern matching only"
    echo ""
    echo "To use AI mode, run:"
    echo "  export OPENAI_API_KEY='your-key-here'"
    echo "  ./demo.sh"
    echo ""
    echo "Get a key from: https://platform.openai.com/api-keys"
    echo ""
    echo "========================================="
    echo ""
fi

# Run the solver
echo "$PUZZLE" | ./target/local/bin/connections
