#!/bin/bash

echo "Testing Gemini API Models"
echo "=========================="
echo ""

API_KEY="AIzaSyBXQZ-BxrFKQMONLvkERpXfRXCUquWk5aU"

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
