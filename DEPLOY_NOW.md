# DEPLOYMENT INSTRUCTIONS - Personal Computer

You've pulled the latest code from GitHub. Here's what to do next:

## ✅ What's Already Done (On Work Computer)
- Fixed Go version in go.mod
- Updated to Gemini 2.5 Flash model
- Created Procfile for Heroku
- Generated go.sum
- Pushed everything to GitHub

## 🚀 Deploy to Heroku (Do This Now)

### Step 1: Set Up Heroku App (First Time Only)

If you already have a Heroku app for this project:
```bash
cd ~/path/to/connections
heroku git:remote -a YOUR-HEROKU-APP-NAME
```

If you need to create a new app:
```bash
cd ~/path/to/connections
heroku create connections-solver-sjg
```

### Step 2: Set Environment Variable
```bash
heroku config:set GEMINI_API_KEY=AIzaSyBXQZ-BxrFKQMONLvkERpXfRXCUquWk5aU
```

### Step 3: Deploy
```bash
git push heroku main
```

Heroku will:
- Detect it's a Go app
- Build the binary from `cmd/cli`
- Run it using the Procfile: `bin/cli`

### Step 4: Check It Worked
```bash
# View logs
heroku logs --tail

# Open the app
heroku open
```

## 📝 Files Created for Heroku

- **Procfile**: Tells Heroku to run `bin/cli` (the built binary)
- **go.mod**: Specifies Go version (1.21)
- **go.sum**: Dependencies checksums (required by Heroku)

## 🔄 Future Updates Workflow

### On Work Computer:
1. Make code changes
2. `git add -A && git commit -m "message"`
3. `git push origin main`

### On Personal Computer:
1. `git pull origin main`
2. `git push heroku main`

That's it! Heroku automatically rebuilds and deploys.

## 🐛 Troubleshooting

### Check your Heroku apps:
```bash
heroku apps
```

### Check environment variables:
```bash
heroku config
```

### View real-time logs:
```bash
heroku logs --tail
```

### Restart the app:
```bash
heroku restart
```

### Check build logs:
```bash
heroku logs --tail | grep "Build"
```

## ✅ What Should Happen

When you deploy:
1. Heroku detects Go
2. Runs `go build` on `cmd/cli`
3. Creates binary at `bin/cli`
4. Runs it with your GEMINI_API_KEY
5. Your Connections solver is live! 🎉

Ready to deploy! Run the commands above on your personal computer.
