# 🚨 API KEY LEAKED - ACTION REQUIRED

## What Happened
Your Gemini API key was leaked and Google has disabled it for security.

**Error message:**
```
AI analysis failed (gemini API error: Your API key was reported as leaked. 
Please use another API key.), falling back to pattern matching...
```

## How It Likely Happened
The API key was accidentally committed to git history or shown in:
- Shell scripts (test-gemini-models.sh, etc.)
- Documentation files
- Commit messages
- GitHub push protection caught it

## What To Do NOW

### Step 1: Create a New API Key
1. Go to Google AI Studio: https://aistudio.google.com/app/apikey
2. Click "Create API Key"
3. Copy the new key (starts with `AIza...`)

### Step 2: Update Your .env File
On your **work computer**:
```bash
cd ~/repos/sjg/connections
echo "GEMINI_API_KEY=YOUR-NEW-KEY-HERE" > .env
```

On your **personal computer** (after you create the key):
```bash
cd ~/path/to/connections
echo "GEMINI_API_KEY=YOUR-NEW-KEY-HERE" > .env
```

### Step 3: Verify .env is NOT in Git
```bash
git status
# Should NOT show .env file
```

The `.env` file is already in `.gitignore` so it won't be committed.

### Step 4: Set on Heroku (Personal Computer)
```bash
heroku config:set GEMINI_API_KEY=YOUR-NEW-KEY-HERE
```

## Files to Check and Clean

I found potential leaks in these files - let me remove them:

### Files with hardcoded keys:
- `test-gemini-models.sh` - Has the old key hardcoded
- `build-and-install.sh` - Might have test keys
- `full-test.sh` - Might have test keys
- `solve` - Might have test keys
- `test-claude.sh` - Might have test keys

These need to be cleaned up and should read from .env instead.

## Security Best Practices Going Forward

### ✅ DO:
- Store API keys ONLY in `.env` file
- `.env` file is in `.gitignore` (already done ✅)
- Use environment variables in code
- Set keys on Heroku with `heroku config:set`

### ❌ DON'T:
- Hardcode API keys in any files
- Commit `.env` to git
- Share API keys in chat/email
- Put keys in documentation or examples

## Testing After New Key

```bash
# On work computer
cd ~/repos/sjg/connections
echo "GEMINI_API_KEY=your-new-key" > .env
go run ./cmd/web
# Open http://localhost:8080 and test
```

## What I'm Fixing Now

I'm going to remove all hardcoded API keys from the repository files
and update them to use environment variables only.
