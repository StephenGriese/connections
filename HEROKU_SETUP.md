# Heroku Deployment Guide

## Fix: "missing required flag app" Error

When you see "missing required flag app", you need to either:
1. Add the app name to the command
2. OR add a Heroku remote to your git repo

## Option 1: Use App Name in Command (Quick Fix)

```bash
# Replace YOUR-APP-NAME with your actual Heroku app name
heroku config:set GEMINI_API_KEY=AIzaSyBXQZ-BxrFKQMONLvkERpXfRXCUquWk5aU --app YOUR-APP-NAME
```

To find your app name:
```bash
heroku apps
```

## Option 2: Add Heroku Remote (Better - Only Do Once)

```bash
# Replace YOUR-APP-NAME with your actual Heroku app name
heroku git:remote -a YOUR-APP-NAME

# Now you can use commands without --app flag
heroku config:set GEMINI_API_KEY=AIzaSyBXQZ-BxrFKQMONLvkERpXfRXCUquWk5aU
```

## If You Don't Have a Heroku App Yet

Create one:
```bash
# Create a new Heroku app
heroku create connections-solver-sjg

# Or let Heroku pick a name:
heroku create

# This will automatically add the heroku remote
```

## Complete Deployment Steps

### 1. Create/Link Heroku App
```bash
# If you already have an app:
heroku git:remote -a YOUR-APP-NAME

# OR create a new one:
heroku create connections-solver-sjg
```

### 2. Set Environment Variable
```bash
heroku config:set GEMINI_API_KEY=AIzaSyBXQZ-BxrFKQMONLvkERpXfRXCUquWk5aU
```

### 3. Create Procfile (if not exists)
```bash
echo "web: ./target/local/bin/connections" > Procfile
git add Procfile
git commit -m "Add Procfile for Heroku"
```

### 4. Deploy
```bash
git push heroku main
```

### 5. Check Logs
```bash
heroku logs --tail
```

### 6. Open App
```bash
heroku open
```

## Verify Configuration

Check your environment variables:
```bash
heroku config
```

Should show:
```
GEMINI_API_KEY: AIzaSy...
```

## Troubleshooting

### List your apps:
```bash
heroku apps
```

### Check which app is linked:
```bash
git remote -v
```

Should show something like:
```
heroku  https://git.heroku.com/your-app-name.git (fetch)
heroku  https://git.heroku.com/your-app-name.git (push)
```

### Remove and re-add remote:
```bash
git remote remove heroku
heroku git:remote -a YOUR-APP-NAME
```

## Quick Reference

```bash
# Set config with app name
heroku config:set KEY=value --app app-name

# Set config with remote (after heroku git:remote)
heroku config:set KEY=value

# View config
heroku config --app app-name
# OR (with remote)
heroku config

# Deploy
git push heroku main

# Logs
heroku logs --tail --app app-name
# OR
heroku logs --tail
```
