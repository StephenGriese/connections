#!/bin/bash

echo "Testing Gemini API Models"
echo "=========================="
echo ""

# Get API key from environment or .env file
if [ -z "$GEMINI_API_KEY" ]; then
    if [ -f ".env" ]; then
        export $(grep -v '^#' .env | xargs)
    fi
fi

if [ -z "$GEMINI_API_KEY" ]; then
    echo "Error: GEMINI_API_KEY not set"
    echo "Set it with: export GEMINI_API_KEY=your-key"
    echo "Or create a .env file with: GEMINI_API_KEY=your-key"
    exit 1
fi

API_KEY="$GEMINI_API_KEY"

# Test different model names
models=(
    "gemini-1.5-flash-002"
    "gemini-1.5-flash"
    "gemini-1.5-flash-latest"
    "gemini-pro"
    "gemini-1.0-pro"
)

for model in "${models[@]}"; do
    echo "Testing: $model"
    response=$(curl -s "https://generativelanguage.googleapis.com/v1beta/models/${model}:generateContent?key=${API_KEY}" \
      -H 'Content-Type: application/json' \
      -d '{"contents":[{"parts":[{"text":"Hi"}]}]}')

    if echo "$response" | grep -q '"candidates"'; then
        echo "  ✅ WORKS!"
    elif echo "$response" | grep -q '"error"'; then
        error_msg=$(echo "$response" | grep -o '"message":"[^"]*"' | head -1)
        echo "  ❌ Error: $error_msg"
    else
        echo "  ❓ Unknown response"
    fi
    echo ""
done
