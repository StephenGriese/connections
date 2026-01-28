# Debugging .env Loading in GoLand

## Breakpoint Location

**File:** `cmd/web/main.go`
**Line:** 38 (marked with 🔴 BREAKPOINT HERE comment)

```go
// 🔴 BREAKPOINT HERE - Line 38: Set breakpoint to see .env loading
if err := godotenv.Load(); err != nil {
```

## How to Set Breakpoint in GoLand

### Method 1: Click in Gutter
1. Open `cmd/web/main.go`
2. Find line 38 (the `if err := godotenv.Load()` line)
3. Click in the left gutter (margin) next to line 38
4. A red dot 🔴 will appear

### Method 2: Keyboard Shortcut
1. Put cursor on line 38
2. Press `Cmd+F8` (Mac) or `Ctrl+F8` (Windows/Linux)

## Run with Debugger

1. Select "Web Server" run configuration
2. Click the **Debug** button 🐞 (next to the green play button)
3. The debugger will stop at line 38

## What You'll See at the Breakpoint

When the debugger stops, you can inspect:

### Variables Panel
- `err` - Will be `nil` if .env loaded, or an error if not found
- After stepping over, check `apiKey` variable

### Evaluate Expression (Alt+F8)
```go
os.Getenv("GEMINI_API_KEY")  // See the actual value
godotenv.Load()               // See the error if any
```

### Console Output
You'll see one of:
```
No .env file found (this is OK on Heroku): ...
```
OR
```
✅ Loaded .env file
   GEMINI_API_KEY loaded: AIzaSyBXQZ-BxrFKQMO...
```

## Step Through the Code

After breakpoint hits:
1. **Step Over** (F8) - Execute line and move to next
2. **Step Into** (F7) - Go inside godotenv.Load() to see how it works
3. **Resume** (F9) - Continue execution

## Check Variables

In the Variables panel, expand to see:
- `err` - The error from godotenv.Load()
- Environment variables after loading

## Verify API Key Loaded

After line 42 executes, check:
```go
apiKey := os.Getenv("GEMINI_API_KEY")
```

Should show your API key starting with "AIza..."

## Console Logging (No Breakpoint Needed)

Even without a breakpoint, when you run the app you'll see:
```
✅ Loaded .env file
   GEMINI_API_KEY loaded: AIzaSyBXQZ-BxrFKQMO...
🚀 Connections Solver API starting on :8080
```

This confirms the .env file was loaded successfully!

## Troubleshooting

### Breakpoint Not Hitting?
- Make sure you clicked the Debug button 🐞 (not the Run button ▶️)
- Verify the red dot is visible in the gutter
- Check you're running "Web Server" configuration

### API Key Empty?
If you see `⚠️ GEMINI_API_KEY is empty!`:
1. Check .env file exists: `ls -la .env`
2. Check .env content: `cat .env`
3. Make sure it says: `GEMINI_API_KEY=your-key-here`
4. No quotes, no spaces around the `=`

### Can't See Variables?
- Switch to "Debug" tab at bottom of GoLand
- Click "Variables" panel
- Expand objects to see their contents
