# DEPLOYMENT INSTRUCTIONS - Personal Computer

You've pulled the latest code from GitHub. Here's what to do next:

## ✅ What's Already Done (On Work Computer)
- Fixed Go version in go.mod
- Updated to Gemini 2.5 Flash model
- Created Procfile for Heroku
- Generated go.sum
- **Created web server wrapper** (Heroku needs HTTP server, not CLI)
- Pushed everything to GitHub

## 🚀 Deploy to Heroku (Do This Now)

### Step 1: Create Your Heroku App (First Time)

Since this is your first deployment, create a new Heroku app:

```bash
cd ~/path/to/connections

# Let Heroku generate a random name (easiest):
heroku create

# OR choose your own name:
heroku create connections-solver-sjg

# OR any name you want (must be unique across all Heroku):
heroku create your-custom-name
```

**What happens:**
- Heroku creates the app
- Automatically adds the `heroku` remote to your git repo
- Gives you a URL like: `https://your-app-name.herokuapp.com`

**Note:** You can add a custom subdomain later in the Heroku dashboard!

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
- Build the web server from `cmd/web`
- Run it using the Procfile: `bin/web`
- Your app will be accessible via web browser!

### Step 4: Check It Worked
```bash
# View logs
heroku logs --tail

# Open the app
heroku open
```

## ⚠️ GOT AN ERROR? - Read This First!

If you see "Application Error" when you open the app, **view the logs** to see what went wrong:

```bash
# View real-time logs (most helpful):
heroku logs --tail

# Or see last 100 lines:
heroku logs -n 100
```

**Fixed:** The app is now a web server that Heroku can run! The H10 errors should be gone.

## 📝 Files Created for Heroku

- **Procfile**: Tells Heroku to run `bin/web` (the web server)
- **cmd/web/main.go**: Web server with HTML interface and API endpoints
- **go.mod**: Specifies Go version (1.21)
- **go.sum**: Dependencies checksums (required by Heroku)
- **cmd/cli/main.go**: Original CLI tool (still works locally)

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
2. Runs `go build` on `cmd/web` (the web server)
3. Creates binary at `bin/web`
4. Starts the web server listening on Heroku's PORT
5. Opens a nice web interface where you can paste words and get solutions! 🎉

**After deployment:** Just open the URL in your browser, paste in 16 words, and click "Solve Puzzle"!

## 🌐 Adding a Custom Domain Later (Optional)

After your app is deployed, you can add a custom subdomain:

1. Go to: https://dashboard.heroku.com/apps/your-app-name/settings
2. Scroll to "Domains"
3. Click "Add domain"
4. Enter your custom domain/subdomain
5. Update your DNS settings as Heroku instructs

For now, your app will be at: `https://your-app-name.herokuapp.com`

Ready to deploy! Run the commands above on your personal computer.
