# GoLand Run Configurations

## Available Configurations

I've created two run configurations for you in GoLand:

### 1. Web Server
**What it does:** Runs the web server on http://localhost:8080

**How to use:**
1. In GoLand, look for the run configuration dropdown (top right)
2. Select "Web Server"
3. Click the green play button ▶️
4. Open your browser to http://localhost:8080
5. Paste in 16 words and solve!

**Environment Variables:**
- `GEMINI_API_KEY` - Set this to your Gemini API key
- `PORT` - Default is 8080

**To set your API key:**
1. Click the run configuration dropdown
2. Select "Edit Configurations..."
3. Find "Web Server"
4. In the "Environment variables" field, set: `GEMINI_API_KEY=your-key-here`
5. Click OK

### 2. CLI (local)
**What it does:** Runs the command-line interface version

**How to use:**
1. Select "CLI (local)" from the run configuration dropdown
2. Click the green play button ▶️
3. Type or paste words in the Run panel
4. See results in the console

**Environment Variables:**
- `GEMINI_API_KEY` - Set this to your Gemini API key

## Quick Setup

### Option 1: Use .env file (Recommended)
Create a `.env` file in the project root:
```bash
GEMINI_API_KEY=AIzaSy...your-key-here
```

Both configurations will automatically load this file!

### Option 2: Set in Run Configuration
1. Run → Edit Configurations...
2. Select the configuration
3. Set environment variables in the UI

## Files Created

```
.idea/runConfigurations/
├── Web_Server.xml      # Web server configuration
└── CLI__local_.xml     # CLI configuration
```

These files are in `.gitignore` so your API key won't be committed!

## Troubleshooting

**Configuration not showing up?**
- Restart GoLand
- Check that the files are in `.idea/runConfigurations/`

**API key not working?**
- Make sure you've set `GEMINI_API_KEY` in the run configuration
- Or create a `.env` file in the project root

**Port already in use?**
- Change the PORT environment variable to 8081 or another port
- Or kill the process using port 8080

## Running from Terminal

You can also run these directly:

```bash
# Web server
GEMINI_API_KEY=your-key PORT=8080 go run ./cmd/web

# CLI
GEMINI_API_KEY=your-key go run ./cmd/cli
```

Enjoy your Connections solver! 🎉
