#!/bin/bash

echo "🔨 Building Connections Solver..."
echo ""

cd ~/repos/sjg/connections || exit 1

# Clean old build
echo "Cleaning old build..."
rm -rf target
mkdir -p target/local/bin

# Build the project
echo "Compiling..."
go build -o target/local/bin/connections ./cmd/cli

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo ""

    # Update ~/bin/connections
    echo "📦 Installing to ~/bin/connections..."
    cat > ~/bin/connections << 'EOFSCRIPT'
#!/bin/bash
cd ~/repos/sjg/connections || exit 1
./target/local/bin/connections
EOFSCRIPT

    chmod +x ~/bin/connections

    echo "✅ Installed!"
    echo ""
    echo "🧪 Testing..."

    # Quick test
    if [ -n "$ANTHROPIC_API_KEY" ]; then
        echo "BASS CLUB DIAMOND HEART" | timeout 5 ./target/local/bin/connections 2>&1 | head -10 || true
    else
        echo "Note: Set ANTHROPIC_API_KEY environment variable to use Claude AI"
    fi

    echo ""
    echo "================================"
    echo "✅ BUILD COMPLETE!"
    echo "================================"
    echo ""
    echo "Run the solver with:"
    echo "  connections"
    echo ""
    echo "Or directly:"
    echo "  ~/bin/connections"
    echo ""
else
    echo "❌ Build failed!"
    echo ""
    echo "Check the error messages above."
    exit 1
fi
