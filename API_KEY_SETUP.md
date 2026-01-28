# API Key Setup Guide

## ⚠️ IMPORTANT: Security First

The old API key `sk-ant-api03-K6TZO...` was exposed in git history and must be revoked immediately.

1. Go to: https://console.anthropic.com/settings/keys
2. Delete the exposed key

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

### Google Gemini (Recommended)
- **Free Tier:** Very generous
- **Get Key:** https://makersuite.google.com/app/apikey
- **Add to .env:** `GEMINI_API_KEY=your-key-here`

### Anthropic Claude
- **Free Tier:** Limited but available
- **Get Key:** https://console.anthropic.com/settings/keys
- **Add to .env:** `ANTHROPIC_API_KEY=your-key-here`

### OpenAI
- **Free Tier:** Trial credits
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
