# Deployment Workflow - Work Computer → Personal Computer → Heroku

## Current Status
✅ All code is committed and pushed to GitHub: `git@github.com-sjg:StephenGriese/connections.git`
✅ Using Gemini 2.5 Flash AI model (working!)
✅ API key stored in `.env` file (gitignored - safe)

## Workflow: Work → Personal → Heroku

### Step 1: On Work Computer (DONE ✅)
```bash
# Everything is already pushed!
cd ~/repos/sjg/connections
git status  # Should show "nothing to commit, working tree clean"
```

### Step 2: On Personal Computer (DO THIS NEXT)
```bash
# Pull the latest changes
cd ~/path/to/connections  # wherever you cloned it on personal computer
git pull origin main

# Rebuild with the new Gemini model
make build

# Test it locally
./target/local/bin/connections
```

### Step 3: Deploy to Heroku (AFTER TESTING)
```bash
# Make sure you're on your personal computer where Heroku is configured
git push heroku main

# Or if you need to force push:
git push heroku main --force
```

## Important Notes

### API Key Management
- **Work Computer:** API key is in `.env` (not committed to git)
- **Personal Computer:** You'll need to add the same API key to `.env` there
- **Heroku:** Set as environment variable (not in code):
  ```bash
  heroku config:set GEMINI_API_KEY=your-key-here
  ```

### The `.env` File
Your `.env` file should have:
```
GEMINI_API_KEY=AIzaSy...your-key-here
```

This file is in `.gitignore` so it won't be committed. You'll need to create it on:
1. ✅ Work computer (already done)
2. ⏳ Personal computer (create manually)
3. ⏳ Heroku (use `heroku config:set`)

## Quick Commands

### Check what's in GitHub:
```bash
git log origin/main --oneline -10
```

### Verify Heroku app:
```bash
heroku apps  # List your apps
heroku config  # See environment variables
heroku logs --tail  # Watch live logs
```

### Test the deployment:
```bash
# After deploying to Heroku, test it:
heroku open
# Or
curl https://your-app-name.herokuapp.com/
```

## What Changed Today
- ✅ Fixed Gemini API to use `gemini-2.5-flash` (was using outdated models)
- ✅ Removed hardcoded API keys from all files
- ✅ Added automatic `.env` file loading
- ✅ Fixed all linting errors
- ✅ Pushed clean code to GitHub

Your next step: Pull on your personal computer and deploy to Heroku! 🚀
