# API Key Setup Guide

## ⚠️ IMPORTANT: Security First

The old API key `sk-ant-api03-K6TZO...` was exposed in git history and has been revoked. ✅

## 💡 Recommended: Use Google Gemini

**Anthropic Claude requires credits/payment**, but Google Gemini has a very generous free tier. 
We recommend using Gemini for this project.

## Adding Your API Key

### Option 1: Use .env file (RECOMMENDED - Safe for Git)

The project now automatically loads from a `.env` file.

1. Edit `~/repos/sjg/connections/.env`
2. Add your key (choose one):
   ```bash
   # Google Gemini (recommended - generous free tier)
   GEMINI_API_KEY=your-key-here
   
   # OR Anthropic Claude
   ANTHROPIC_API_KEY=your-key-here
   
   # OR OpenAI
   OPENAI_API_KEY=your-key-here
   ```
3. The `.env` file is already in `.gitignore`, so it won't be committed

### Option 2: Use ~/.zshrc (System-wide)

Add to your shell configuration:

```bash
# Add to ~/.zshrc
echo 'export ANTHROPIC_API_KEY="your-new-key-here"' >> ~/.zshrc

# Reload shell
source ~/.zshrc
```

**Note:** Be careful not to commit files with this environment variable set!

## Getting API Keys

### Google Gemini (⭐ Recommended)
- **Free Tier:** Very generous - 1,500 requests per day!
- **Cost:** FREE for personal use
- **Get Key:** https://makersuite.google.com/app/apikey (or https://aistudio.google.com/app/apikey)
- **Add to .env:** `GEMINI_API_KEY=your-key-here`

### Anthropic Claude
- **Free Tier:** ❌ Not available - requires payment/credits
- **Cost:** Pay-as-you-go (credit balance required)
- **Get Key:** https://console.anthropic.com/settings/keys
- **Add to .env:** `ANTHROPIC_API_KEY=your-key-here`
- **Note:** You'll see errors if your credit balance is too low

### OpenAI
- **Free Tier:** Trial credits (limited)
- **Cost:** Pay-as-you-go after trial
- **Get Key:** https://platform.openai.com/api-keys
- **Add to .env:** `OPENAI_API_KEY=your-key-here`

## Priority Order

The app will try APIs in this order:
1. Google Gemini (if `GEMINI_API_KEY` is set)
2. Anthropic Claude (if `ANTHROPIC_API_KEY` is set)
3. OpenAI (if `OPENAI_API_KEY` is set)
4. Pattern matching (if no API keys)

## Testing

```bash
cd ~/repos/sjg/connections
make build
./target/local/bin/connections
```

Or if installed in `~/bin`:
```bash
connections
```

## How It Works

The app automatically:
1. Looks for `.env` file in the project directory
2. Loads any API keys found there
3. Falls back to environment variables if set
4. Uses pattern matching if no API keys are available

## Troubleshooting

### Gemini API Error: "model is not found for API version"

The app now uses:
- API version: `v1beta`
- Model: `gemini-2.5-flash` (current stable model as of 2025)

**Note:** Google updated their models. Old models like `gemini-1.5-flash` and `gemini-pro` 
no longer exist. Current models include `gemini-2.5-flash`, `gemini-2.5-pro`, etc.

Make sure to rebuild after updating:
```bash
cd ~/repos/sjg/connections
make build
```

### Verify Your Gemini API Key Works

Test your API key directly with curl:

```bash
curl "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=YOUR_API_KEY" \
  -H 'Content-Type: application/json' \
  -d '{"contents":[{"parts":[{"text":"Hello"}]}]}'
```

Replace `YOUR_API_KEY` with your actual Gemini API key. You should see a response with 
`"candidates"` containing generated text.

### Claude API Error: "credit balance is too low"

Anthropic Claude requires payment/credits. Switch to Google Gemini (free) instead.

### No API Key Found

Make sure your `.env` file has the key on its own line:
```
GEMINI_API_KEY=AIzaSy...your-actual-key
```

No quotes needed, no extra spaces.

