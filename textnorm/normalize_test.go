package textnorm

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNormalizeCases(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile(filepath.Join("..", "testdata", "cases.json"))
	if err != nil {
		t.Fatalf("read cases: %v", err)
	}

	var cases []struct {
		Name  string `json:"name"`
		Input string `json:"input"`
		Want  string `json:"want"`
	}
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("decode cases: %v", err)
	}

	for _, testCase := range cases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			if got := Normalize(testCase.Input); got != testCase.Want {
				t.Fatalf("Normalize(%q) = %q, want %q", testCase.Input, got, testCase.Want)
			}
		})
	}
}
