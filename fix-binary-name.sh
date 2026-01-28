#!/bin/bash

echo "🔧 Fixing binary name from 'wordle' to 'connections'..."
echo ""

cd ~/repos/sjg/connections || exit 1

# Remove old binary
echo "1. Removing old 'wordle' binary..."
rm -rf target

# Build with correct name
echo "2. Building 'connections' binary..."
make build

# Verify
if [ -f target/local/bin/connections ]; then
    echo ""
    echo "✅ SUCCESS! Binary built:"
    ls -lh target/local/bin/connections
    echo ""
    echo "3. Testing the binary..."
    echo "   Run: connections"
    echo ""
else
    echo ""
    echo "❌ Build may have failed. Checking what exists..."
    ls -la target/local/bin/ 2>&1 || echo "No bin directory found"
    echo ""
    echo "Trying manual build..."
    mkdir -p target/local/bin
    go build -o target/local/bin/connections ./cmd/cli

    if [ -f target/local/bin/connections ]; then
        echo "✅ Manual build succeeded!"
    else
        echo "❌ Manual build also failed. Check for errors above."
    fi
fi

echo ""
echo "Done!"
