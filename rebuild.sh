#!/bin/bash

echo "🔨 Rebuilding the solver..."
cd ~/repos/sjg/connections

# Clean and rebuild
rm -f target/local/bin/connections
mkdir -p target/local/bin
go build -o target/local/bin/connections ./cmd/cli

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo ""
    echo "🧪 Quick test..."
    echo ""

    # Test if API key is available in environment
    if [ -n "$ANTHROPIC_API_KEY" ] || [ -n "$GEMINI_API_KEY" ] || [ -n "$OPENAI_API_KEY" ]; then
        echo "BASS CLUB DIAMOND HEART SPADE SOLE PERCH TROUT WOOD IRON DRIVER PUTTER ACE KING QUEEN JACK" | ./target/local/bin/connections
    else
        echo "Note: Set ANTHROPIC_API_KEY, GEMINI_API_KEY, or OPENAI_API_KEY for AI mode"
    fi

    echo ""
    echo "🔄 Updating ~/bin/connections..."
    cp target/local/bin/connections ~/bin/connections-binary

    # Update the wrapper script
    cat > ~/bin/connections << 'EOF'
#!/bin/bash
# NYTimes Connections Solver
cd ~/repos/sjg/connections || exit 1
./target/local/bin/connections
EOF

    chmod +x ~/bin/connections

    echo "✅ Done! Try running: connections"
else
    echo "❌ Build failed!"
    exit 1
fi
