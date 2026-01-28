# ⚠️ IMPORTANT: Environment Variable Precedence Issue

## What Happened
If you set `GEMINI_API_KEY` in GoLand's run configuration environment variables, it will **OVERRIDE** the `.env` file!

This is a common gotcha that can cause confusion.

## Environment Variable Priority

When you run the app, environment variables are loaded in this order (highest priority first):

1. **🔴 Run Configuration env vars** ← Set in "Edit Configurations"  
   **Overrides everything below!**

2. **Shell/System env vars** ← Set in terminal with `export`

3. **✅ .env file** ← Auto-loaded by the app  
   **Recommended location for API keys**

## The Problem This Causes

**Scenario:**
1. You put an old API key in "Edit Configurations" → Environment variables
2. Later, you create a new key and put it in `.env`
3. You run the app... **but it still uses the OLD key!**
4. You get "API key was leaked" error

**Why?** The run configuration env var (step 1) has higher priority than `.env` (step 3).

## The Solution

### ✅ DO THIS:
1. Put your API key ONLY in `.env` file
2. Leave run configuration env vars EMPTY (or remove `GEMINI_API_KEY` line entirely)
3. The app auto-loads from `.env` - you'll see: `✅ Loaded .env file`

### ❌ DON'T DO THIS:
- Don't set `GEMINI_API_KEY` in "Edit Configurations" → Environment variables
- It will override your `.env` file and cause confusion

## How to Check Your Run Configuration

1. Click run configuration dropdown (top right)
2. Select "Edit Configurations..."
3. Select "Web Server" or "CLI (local)"
4. Look at "Environment variables" section
5. **Remove `GEMINI_API_KEY` if it's there!**
6. Click OK

## How to Verify It's Working

Run the app and check the console:

**Good - Using .env file:**
```
✅ Loaded .env file
   GEMINI_API_KEY loaded: AIzaSyBXQZ-BxrFKQMO...
🚀 Connections Solver API starting on :8080
```

**Bad - .env file ignored:**
```
No .env file found
# OR
✅ Loaded .env file
   ⚠️  GEMINI_API_KEY is empty!
```

If you see "empty", it means the run configuration env var (set to empty string) is overriding your `.env` file!

## Quick Fix

**Option 1: Remove from Run Configuration (Recommended)**
1. Edit Configurations → Web Server
2. Environment variables: Remove or clear `GEMINI_API_KEY`
3. Leave it blank - the app will use `.env`

**Option 2: Update Run Configuration Value**
1. Edit Configurations → Web Server  
2. Environment variables: Set `GEMINI_API_KEY=your-new-key-here`
3. But this means you have to update it in TWO places when you rotate keys!

**Option 1 is better!** Single source of truth in `.env` file.

## Remember
- `.env` file is in `.gitignore` ✅ (won't be committed)
- Run configuration files ARE committed to git
- So using `.env` is more secure anyway!
