package ai

import (
	"testing"
)

func TestParseJSONResponse(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name: "valid JSON response",
			content: `[
				{
					"words": ["BASS", "TROUT", "PERCH", "SOLE"],
					"theme": "Fish",
					"explanation": "All are types of fish",
					"confidence": 0.95
				},
				{
					"words": ["CLUB", "DIAMOND", "HEART", "SPADE"],
					"theme": "Card Suits",
					"explanation": "Playing card suits",
					"confidence": 0.98
				},
				{
					"words": ["WOOD", "IRON", "DRIVER", "PUTTER"],
					"theme": "Golf Clubs",
					"explanation": "Types of golf clubs",
					"confidence": 0.92
				},
				{
					"words": ["ACE", "KING", "QUEEN", "JACK"],
					"theme": "Face Cards",
					"explanation": "Face cards and ace",
					"confidence": 0.90
				}
			]`,
			wantErr: false,
		},
		{
			name: "JSON in markdown code block",
			content: "```json\n" + `[
				{
					"words": ["BASS", "TROUT", "PERCH", "SOLE"],
					"theme": "Fish",
					"explanation": "All are types of fish",
					"confidence": 0.95
				},
				{
					"words": ["CLUB", "DIAMOND", "HEART", "SPADE"],
					"theme": "Card Suits",
					"explanation": "Playing card suits",
					"confidence": 0.98
				},
				{
					"words": ["WOOD", "IRON", "DRIVER", "PUTTER"],
					"theme": "Golf Clubs",
					"explanation": "Types of golf clubs",
					"confidence": 0.92
				},
				{
					"words": ["ACE", "KING", "QUEEN", "JACK"],
					"theme": "Face Cards",
					"explanation": "Face cards and ace",
					"confidence": 0.90
				}
			]` + "\n```",
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			content: "This is not JSON",
			wantErr: true,
		},
		{
			name: "partial results - 1 group (now accepted)",
			content: `[
				{
					"words": ["BASS", "TROUT", "PERCH", "SOLE"],
					"theme": "Fish",
					"explanation": "All are types of fish",
					"confidence": 0.95
				}
			]`,
			wantErr: false, // Changed: we now accept partial results
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			groups, err := parseJSONResponse(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseJSONResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// We now accept 1-4 groups (partial results)
				if len(groups) < 1 || len(groups) > 4 {
					t.Errorf("parseJSONResponse() returned %d groups, want 1-4", len(groups))
				}
			}
		})
	}
}
