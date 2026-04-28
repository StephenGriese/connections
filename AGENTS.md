# AGENTS.md — connections

NYT Connections AI solver with both a CLI and a web interface.

## Commands

```bash
go build ./...       # build all packages
go test ./...        # run all tests
```

## Package layout

- `cmd/cli/` — CLI entrypoint
- `cmd/web/` — Web server entrypoint
- `pkg/ai/` — AI provider implementations (OpenAI, Claude, Gemini)

## HTTP client conventions

See the workspace-level `../AGENTS.md` for the full rules. Key points for this project:

- Each AI provider (`OpenAIProvider`, `ClaudeProvider`, `GeminiProvider`) stores its
  `*http.Client` as a struct field initialized in the constructor with a 60s timeout.
- Do not create `&http.Client{}` inline inside `AnalyzeWords` or any other method.
- Use `http.NewRequestWithContext` for all outbound requests.
- Base URLs should be named constants, not inline strings.

## AI provider timeout

All three providers use `defaultAITimeout = 60 * time.Second`. Adjust if an API is
consistently slower, but never use a bare `&http.Client{}` with no timeout.
