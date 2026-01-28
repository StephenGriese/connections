#!/bin/bash

echo "================================"
echo "COMPREHENSIVE BUILD & TEST"
echo "================================"
echo ""

cd ~/repos/sjg/connections || exit 1

echo "1. Cleaning..."
rm -rf target
mkdir -p target/local/bin

echo ""
echo "2. Testing individual packages..."
echo ""

echo "   Testing pkg/analyzer..."
go test ./pkg/analyzer
if [ $? -ne 0 ]; then
    echo "   ❌ pkg/analyzer FAILED"
else
    echo "   ✅ pkg/analyzer PASSED"
fi

echo ""
echo "   Testing pkg/ai..."
go test ./pkg/ai
if [ $? -ne 0 ]; then
    echo "   ❌ pkg/ai FAILED"
else
    echo "   ✅ pkg/ai PASSED"
fi

echo ""
echo "   Testing pkg/solver..."
go test ./pkg/solver
if [ $? -ne 0 ]; then
    echo "   ❌ pkg/solver FAILED"
    echo ""
    echo "   Getting detailed error:"
    go test -v ./pkg/solver
else
    echo "   ✅ pkg/solver PASSED"
fi

echo ""
echo "3. Building main binary..."
go build -o target/local/bin/connections ./cmd/cli
if [ $? -ne 0 ]; then
    echo "   ❌ Build FAILED"
    exit 1
else
    echo "   ✅ Build SUCCEEDED"
fi

echo ""
echo "4. Installing to ~/bin..."
cat > ~/bin/connections << 'EOFSCRIPT'
#!/bin/bash
cd ~/repos/sjg/connections || exit 1
./target/local/bin/connections
EOFSCRIPT

chmod +x ~/bin/connections
echo "   ✅ Installed to ~/bin/connections"

echo ""
echo "================================"
echo "BUILD COMPLETE!"
echo "================================"
echo ""
echo "Run with: connections"
echo ""
