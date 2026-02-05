package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"connections/pkg/solver"

	"github.com/joho/godotenv"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

// Request payload for the API
type SolveRequest struct {
	Words []string `json:"words"`
}

// Response payload for the API
type SolveResponse struct {
	Success bool    `json:"success"`
	Groups  []Group `json:"groups,omitempty"`
	Error   string  `json:"error,omitempty"`
}

type Group struct {
	Words       []string `json:"words"`
	Theme       string   `json:"theme"`
	Explanation string   `json:"explanation"`
	Confidence  float64  `json:"confidence"`
}

func main() {
	// Load .env file if it exists (for local development)
	// 🔴 BREAKPOINT HERE - Line 38: Set breakpoint to see .env loading
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found (this is OK on Heroku): %v", err)
	} else {
		log.Printf("✅ Loaded .env file")
		// Show what was loaded (first 20 chars of API key for verification)
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey != "" {
			log.Printf("   GEMINI_API_KEY loaded: %s...", apiKey[:min(20, len(apiKey))])
		} else {
			log.Printf("   ⚠️  GEMINI_API_KEY is empty!")
		}
	}

	// Get port from environment (Heroku provides this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local testing
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/solve", handleSolve)
	http.HandleFunc("/health", handleHealth)

	addr := ":" + port
	log.Printf("🚀 Connections Solver API starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	page := h.HTML(
		h.Lang("en"),
		h.Head(
			g.El("meta", g.Attr("charset", "UTF-8")),
			h.TitleEl(g.Text("NYTimes Connections Solver")),
			h.StyleEl(g.Raw(`
				body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
				h1 { color: #333; }
				textarea { width: 100%; height: 150px; padding: 10px; font-size: 16px; }
				button { background: #4CAF50; color: white; padding: 15px 32px; font-size: 16px; border: none; cursor: pointer; margin-top: 10px; }
				button:hover { background: #45a049; }
				#result { margin-top: 20px; padding: 20px; background: #f5f5f5; border-radius: 5px; }
				.group { margin: 10px 0; padding: 10px; background: white; border-left: 4px solid #4CAF50; }
				.error { color: red; }
				.grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 10px; }
				.grid input { padding: 10px; font-size: 16px; width: 100%; box-sizing: border-box; }
			`)),
		),
		h.Body(
			h.H1(g.Text("🔗 NYTimes Connections Solver")),
			h.P(g.Text("Enter 16 words or phrases (one per box). Multi-word phrases like 'bald eagle' are allowed.")),
			// 4x4 grid of 16 inputs
			h.Div(
				h.Class("grid"),
				// ...existing code... (inputs will be expanded below)
				h.Input(h.ID("w0"), h.Placeholder("1")),
				h.Input(h.ID("w1"), h.Placeholder("2")),
				h.Input(h.ID("w2"), h.Placeholder("3")),
				h.Input(h.ID("w3"), h.Placeholder("4")),
				h.Input(h.ID("w4"), h.Placeholder("5")),
				h.Input(h.ID("w5"), h.Placeholder("6")),
				h.Input(h.ID("w6"), h.Placeholder("7")),
				h.Input(h.ID("w7"), h.Placeholder("8")),
				h.Input(h.ID("w8"), h.Placeholder("9")),
				h.Input(h.ID("w9"), h.Placeholder("10")),
				h.Input(h.ID("w10"), h.Placeholder("11")),
				h.Input(h.ID("w11"), h.Placeholder("12")),
				h.Input(h.ID("w12"), h.Placeholder("13")),
				h.Input(h.ID("w13"), h.Placeholder("14")),
				h.Input(h.ID("w14"), h.Placeholder("15")),
				h.Input(h.ID("w15"), h.Placeholder("16")),
			),
			h.Br(),
			h.Button(
				g.Attr("onclick", "solve()"),
				g.Text("Solve Puzzle"),
			),
			h.Div(h.ID("result")),
			h.Script(g.Raw(`
			async function solve() {
				const words = [];
				for (let i = 0; i < 16; i++) {
					const v = (document.getElementById('w' + i).value || '').trim();
					if (v.length === 0) {
						document.getElementById('result').innerHTML = '<p class="error">Please fill in all 16 boxes. Box ' + (i+1) + ' is empty.</p>';
						return;
					}
					words.push(v);
				}

				if (words.length !== 16) {
					document.getElementById('result').innerHTML = '<p class="error">Please enter exactly 16 words.</p>';
					return;
				}

				document.getElementById('result').innerHTML = '<p>Analyzing with Gemini AI...</p>';

				try {
					const response = await fetch('/solve', {
						method: 'POST',
						headers: { 'Content-Type': 'application/json' },
						body: JSON.stringify({ words: words })
					});

					const data = await response.json();

					if (data.success) {
						let html = '<h2>✅ Found ' + data.groups.length + ' groups:</h2>';
						data.groups.forEach((group, i) => {
							html += '<div class="group">';
							html += '<strong>Group ' + (i+1) + ':</strong> ' + group.theme + '<br>';
							html += '<strong>Words:</strong> ' + group.words.join(', ') + '<br>';
							html += '<strong>Explanation:</strong> ' + group.explanation + '<br>';
							html += '<strong>Confidence:</strong> ' + Math.round(group.confidence * 100) + '%';
							html += '</div>';
						});
						document.getElementById('result').innerHTML = html;
					} else {
						document.getElementById('result').innerHTML = '<p class="error">Error: ' + data.error + '</p>';
					}
				} catch (error) {
					document.getElementById('result').innerHTML = '<p class="error">Error: ' + error.message + '</p>';
				}
			}
		`)),
		),
	)

	_ = page.Render(w)
}

func handleSolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SolveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, SolveResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	if len(req.Words) != 16 {
		respondJSON(w, SolveResponse{
			Success: false,
			Error:   fmt.Sprintf("Expected 16 words, got %d", len(req.Words)),
		})
		return
	}

	// Normalize words (uppercase)
	for i, word := range req.Words {
		req.Words[i] = strings.ToUpper(strings.TrimSpace(word))
	}

	// Solve the puzzle with Gemini AI
	apiKey := os.Getenv("GEMINI_API_KEY")
	var s *solver.Solver
	if apiKey != "" {
		log.Printf("Using Gemini AI with API key: %s...", apiKey[:20])
		s = solver.NewWithGemini(apiKey)
	} else {
		log.Printf("No API key found, using pattern matching")
		s = solver.New()
	}

	groups, err := s.Solve(req.Words)
	if err != nil {
		log.Printf("Solver error: %v (found %d groups)", err, len(groups))

		// If we got some groups but not all 4, return them with a warning
		if len(groups) > 0 {
			respGroups := make([]Group, len(groups))
			for i, grp := range groups {
				respGroups[i] = Group{
					Words:       grp.Words,
					Theme:       grp.Theme,
					Explanation: grp.Explanation,
					Confidence:  grp.Confidence,
				}
			}
			respondJSON(w, SolveResponse{
				Success: false,
				Groups:  respGroups,
				Error:   fmt.Sprintf("Only found %d of 4 groups. Try rephrasing or checking your words.", len(groups)),
			})
			return
		}

		respondJSON(w, SolveResponse{
			Success: false,
			Error:   fmt.Sprintf("Solver failed: %v. Make sure you entered exactly 16 valid words.", err),
		})
		return
	}

	// Convert to response format
	respGroups := make([]Group, len(groups))
	for i, grp := range groups {
		respGroups[i] = Group{
			Words:       grp.Words,
			Theme:       grp.Theme,
			Explanation: grp.Explanation,
			Confidence:  grp.Confidence,
		}
	}

	respondJSON(w, SolveResponse{
		Success: true,
		Groups:  respGroups,
	})
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprint(w, `{"status":"ok"}`)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
