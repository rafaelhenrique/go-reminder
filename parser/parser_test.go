package parser

import "testing"

var testCases = []struct {
	input    []string
	expected string
}{
	{
		[]string{"remind", "me",
			"to", "drink", "water",
			"at", "3pm", "every", "day",
		},
		"me|to drink water|at 3pm every day",
	},
	{
		[]string{"remind", "me",
			"to", "update", "the", "project", "status",
			"every", "Monday", "at", "9am",
		},
		"me|to update the project status|every Monday at 9am",
	},
}

func TestParseArgs(t *testing.T) {
	for _, tt := range testCases {
		actual := parseArgs(tt.input)
		if actual != tt.expected {
			msg := `
            input=%v
            expected=%v
            actual=%v
            `
			t.Fatalf(msg, tt.input, tt.expected, actual)
		}
	}
}
