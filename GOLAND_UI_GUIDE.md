# 🎯 Using Run Configurations in GoLand

## The Files I Created ARE GoLand Configurations!

The `.run/*.xml` files I created are GoLand's native run configuration format. They'll appear automatically in your IDE.

---

## How to Access in GoLand

### Method 1: Run Configurations Dropdown (Top-Right)
1. Look at the **top-right corner** of GoLand
2. You'll see a dropdown (currently might say "Current File" or similar)
3. Click it - you'll see:
   - **Connections Solver** ⭐
   - Connections with Gemini
   - Test All
   - Build
4. Select one and click the green ▶️ button

### Method 2: Edit Configurations Menu
1. Click **Run** menu at the top
2. Select **Edit Configurations...**
3. You'll see a dialog with all your configurations listed on the left:
   ```
   ├── Go Build
   ├── Connections Solver          ← Your configs appear here!
   ├── Connections with Gemini
   ├── Test All
   └── Build
   ```

### Method 3: Keyboard Shortcut
- **macOS**: `Cmd + ,` then select "Run/Debug Configurations"
- **Windows/Linux**: Right-click in editor → "More Run/Debug" → "Modify Run Configuration"

---

## What You'll See in Edit Configurations

When you click **Run → Edit Configurations**, you'll see:

### Left Panel (Configuration List)
```
Application
├── Connections Solver          ← Click to edit
├── Connections with Gemini
Go Test
├── Test All
Makefile
└── Build
```

### Right Panel (Configuration Settings)

When you select "Connections Solver", you'll see these tabs:

#### **Configuration Tab:**
- **Name**: Connections Solver
- **Run kind**: Package
- **Package path**: connections/cmd/cli
- **Working directory**: /path/to/connections
- **Environment variables**: 
  - `ANTHROPIC_API_KEY=sk-ant-api03...`

#### **Environment Tab:**
Where you can add/edit environment variables like:
- `ANTHROPIC_API_KEY` (already set)
- `GEMINI_API_KEY` (add if you want)
- `OPENAI_API_KEY` (add if you want)

---

## How to Edit a Configuration

### To Modify "Connections Solver":

1. **Run → Edit Configurations...**
2. **Click "Connections Solver"** on the left
3. **Modify any of these:**

   **General Settings:**
   - Name: Change if you want
   - Package path: `connections/cmd/cli` ✅
   - Working directory: Your project root ✅

   **Environment Variables:**
   - Click the folder icon next to "Environment variables"
   - You'll see: `ANTHROPIC_API_KEY=sk-ant-...`
   - Click `+` to add more
   - Click `Edit` to change existing ones

   **Before Launch:**
   - Add build tasks
   - Add tests to run first
   - etc.

4. **Click OK** to save

---

## How to Run from GoLand UI

### Quick Run:
1. **Select "Connections Solver"** from dropdown (top-right)
2. **Click green ▶️ Run button** (or press `Ctrl+R`/`Cmd+R`)
3. **Enter words in console** when prompted
4. **See results!**

### Debug Mode:
1. **Select "Connections Solver"**
2. **Click bug icon 🐞** (or press `Ctrl+D`/`Cmd+D`)
3. **Set breakpoints** in your code
4. **Step through** the AI logic!

### Run with Modifications:
1. **Right-click "Connections Solver"** in dropdown
2. **Select "Edit 'Connections Solver'..."**
3. **Make changes** (add env vars, change settings)
4. **Click Run** to use modified config

---

## Adding Your Own Configuration

### Create New Go Application Config:

1. **Run → Edit Configurations...**
2. **Click `+` (top-left)**
3. **Select "Go Build"**
4. **Fill in:**
   - Name: `My Custom Connections`
   - Run kind: `Package`
   - Package: `connections/cmd/cli`
   - Working directory: `$PROJECT_DIR$`
5. **Add Environment Variables:**
   - Click folder icon
   - Add your API keys
6. **Click OK**

---

## Common Customizations

### Add Gemini API Key:
1. Edit "Connections with Gemini" config
2. Environment variables → Edit
3. Find: `GEMINI_API_KEY`
4. Set value: `AIzaYourKeyHere`
5. Save and run!

### Add Program Arguments:
1. Edit configuration
2. Find "Program arguments" field
3. Add arguments (though connections doesn't use any currently)

### Change Output Directory:
1. Edit configuration
2. Find "Output directory"
3. Change from `target/local/bin` to whatever you want

### Run Tests Before Launch:
1. Edit configuration
2. "Before launch" section at bottom
3. Click `+` → "Run Another Configuration"
4. Select "Test All"
5. Now tests run automatically before the app!

---

## Visual Guide

When you open **Run → Edit Configurations**, you'll see:

```
┌─────────────────────────────────────────────────────┐
│ Run/Debug Configurations                            │
├──────────────────┬──────────────────────────────────┤
│ [+] [-] [Copy]   │  Configuration: Connections Sol..│
│                  │                                   │
│ ▼ Go Build       │  Name: Connections Solver        │
│   • Connections  │  Package: connections/cmd/cli    │
│     Solver ←     │  Working dir: $PROJECT_DIR$      │
│   • Connections  │                                   │
│     with Gemini  │  Environment variables:          │
│                  │  ANTHROPIC_API_KEY=sk-ant...     │
│ ▼ Go Test        │  [Edit]                          │
│   • Test All     │                                   │
│                  │  ┌─────────────────────────────┐ │
│ ▼ Makefile       │  │ Before launch:              │ │
│   • Build        │  │ • Build                     │ │
│                  │  └─────────────────────────────┘ │
└──────────────────┴──────────────────────────────────┘
        [OK]  [Cancel]  [Apply]
```

---

## Keyboard Shortcuts (in GoLand)

**Run current configuration**: 
- Mac: `Ctrl + R` or `⌃ + R`
- Windows/Linux: `Shift + F10`

**Debug current configuration**:
- Mac: `Ctrl + D` or `⌃ + D`
- Windows/Linux: `Shift + F9`

**Edit configurations**:
- Mac: From Run menu
- Windows/Linux: From Run menu or `Alt + Shift + F10`

**Quick switch configuration**:
- Mac: `Ctrl + Alt + R`
- Windows/Linux: `Alt + Shift + F10`

---

## Troubleshooting

### "I don't see the configurations!"
1. **Restart GoLand** - It should auto-detect `.run/` files
2. Check that `.run/` folder exists in project root
3. Go to Run → Edit Configurations and click `+` to import if needed

### "Environment variables aren't working!"
1. Edit the configuration
2. Check the "Environment variables" field
3. Make sure the API key is set correctly
4. Try clicking "Edit environment variables" and re-entering

### "It's using the wrong API key!"
Priority order:
1. Configuration environment variables (highest)
2. System environment variables
3. .env file

So the run configuration should override everything.

---

## What's Already Configured

✅ **Connections Solver**: 
- Claude API key set
- Ready to run immediately
- No additional setup needed

✅ **Connections with Gemini**:
- Gemini API key placeholder
- Add your key and it's ready

✅ **Test All**:
- Runs all tests
- No configuration needed

✅ **Build**:
- Runs make build
- Includes linting + tests

---

## Quick Start (Right Now!)

1. **Look at top-right of GoLand window**
2. **Find the dropdown** (might say "Current File")
3. **Click it** - You should see "Connections Solver"
4. **Select "Connections Solver"**
5. **Click the green ▶️ Run button**
6. **Type 16 words** when prompted
7. **See Claude solve it!** ✨

---

## That's It!

The XML files I created ARE the GoLand run configurations. They'll show up automatically in:
- Run configurations dropdown (top-right)
- Run → Edit Configurations menu
- Right-click → Run menu

**No import or manual setup required!** Just select and run. 🚀
