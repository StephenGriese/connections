# ✅ GOLAND RUN CONFIGURATIONS CREATED!

## What I Created

I've added **4 GoLand run configurations** to your project in the `.run/` directory:

### 1. **Connections Solver** (Main App)
- **File**: `.run/Connections Solver.run.xml`
- **Type**: Go Application
- **Environment**: Claude API key pre-configured
- **Use**: Run the solver with AI enabled

### 2. **Connections with Gemini**
- **File**: `.run/Connections with Gemini.run.xml`
- **Type**: Go Application
- **Environment**: Gemini API key (empty - add yours)
- **Use**: Run with Google Gemini instead of Claude

### 3. **Test All**
- **File**: `.run/Test All.run.xml`
- **Type**: Go Test
- **Use**: Run all tests in the project

### 4. **Build**
- **File**: `.run/Build.run.xml`
- **Type**: Makefile Target
- **Use**: Run `make build` from GoLand

## How to Use in GoLand

### Method 1: Run Configurations Dropdown
1. Look at the top-right of GoLand (next to the Run/Debug buttons)
2. Click the dropdown that says "Current File" or similar
3. You'll see:
   - **Connections Solver** ← Use this one!
   - Connections with Gemini
   - Test All
   - Build

4. Select "Connections Solver" and click the green ▶️ Run button

### Method 2: Right-Click Menu
1. In the `.run/` folder, right-click on `Connections Solver.run.xml`
2. Select "Run 'Connections Solver'"

### Method 3: Keyboard Shortcut
1. Select "Connections Solver" from the dropdown
2. Press `Ctrl+R` (or `Cmd+R` on Mac) to run
3. Press `Ctrl+D` (or `Cmd+D` on Mac) to debug

## What Each Configuration Does

### Connections Solver (Main)
```
✅ Uses Claude AI (API key included)
✅ Runs from cmd/cli/main.go
✅ Working directory: project root
✅ Output: target/local/bin
```

### Connections with Gemini
```
🔧 Uses Gemini (add your API key in the config)
✅ Runs from cmd/cli/main.go
✅ Working directory: project root
```

To add your Gemini key:
1. Edit `.run/Connections with Gemini.run.xml`
2. Find: `<env name="GEMINI_API_KEY" value="" />`
3. Add your key: `<env name="GEMINI_API_KEY" value="AIza..." />`

### Test All
```
✅ Runs all tests: go test ./...
✅ Shows test results in GoLand
✅ Can run with coverage
```

### Build
```
✅ Runs: make build
✅ Includes all linting and tests
✅ Compiles the binary
```

## Quick Start

1. **Open GoLand** (restart if already open to pick up configs)
2. **Select "Connections Solver"** from run config dropdown
3. **Click Run** ▶️
4. **Enter 16 words** when prompted
5. **See Claude solve it!** ✨

## Running with Debug

1. Select "Connections Solver"
2. Click Debug 🐞 (or press Cmd+D)
3. Set breakpoints in your code
4. Step through the AI analysis!

## Environment Variables

### Already Configured (Claude):
- `ANTHROPIC_API_KEY` - Your Claude key is set in "Connections Solver"

### To Configure (Gemini):
Edit `.run/Connections with Gemini.run.xml` and add your key.

### To Configure (OpenAI):
Create a new run configuration or edit an existing one:
```xml
<env name="OPENAI_API_KEY" value="sk-your-key" />
```

## Customizing Run Configurations

You can edit any `.run/*.xml` file to:
- Add more environment variables
- Change program arguments
- Modify working directory
- Add before launch tasks

Or use GoLand's UI:
1. Run → Edit Configurations
2. Select a configuration
3. Modify settings
4. Click OK

## Tips

### Run with Input File
To test with pre-defined input:
1. Create a file: `test-input.txt` with 16 words
2. Edit run config → Add program arguments
3. Or pipe in: `cat test-input.txt | connections`

### Debug AI Calls
Set breakpoints in:
- `pkg/ai/provider.go` - See API requests/responses
- `pkg/solver/solver.go` - See solving logic
- `cmd/cli/main.go` - See CLI flow

### Run Tests for One Package
1. Right-click on a `*_test.go` file
2. Select "Run Tests in File"
3. Or use keyboard shortcut

## File Structure

```
.run/
├── Connections Solver.run.xml      ← Main app with Claude
├── Connections with Gemini.run.xml ← Alternative with Gemini
├── Test All.run.xml                ← Run all tests
└── Build.run.xml                   ← Make build
```

## Gitignore Note

The `.run/` directory **should be committed** to git so your team can use these configs too!

If you want to keep API keys private:
1. Use a separate `.env` file (already in .gitignore)
2. Or use GoLand's "Store as project file" option

---

## ✅ READY TO USE!

Your GoLand run configurations are set up and ready to go!

**Next Steps:**
1. Open/restart GoLand
2. Select "Connections Solver" from dropdown
3. Click Run ▶️
4. Solve puzzles! 🎉

The configurations will appear in GoLand automatically. No manual import needed!
