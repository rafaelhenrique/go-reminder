package parser

import "testing"

var testCases = []struct {
	input   []string
	someone string
	what    string
	when    string
}{
	{
		[]string{"remind", "me",
			"to", "drink", "water",
			"at", "3pm", "every", "day",
		},
		"me", "to drink water", "at 3pm every day",
	},
	{
		[]string{"remind", "me",
			"to", "update", "the", "project", "status",
			"every", "Monday", "at", "9am",
		},
		"me", "to update the project status", "every Monday at 9am",
	},
}

func TestParseArgs(t *testing.T) {
	for _, tt := range testCases {
		someone, what, when := ParseArgs(tt.input)

		if someone != tt.someone || what != tt.what || when != tt.when {
			msg := `
            input=%v
            expected=%v<space>%v<space>%v
            actual=%v<space>%v<space>%v
            `
			t.Fatalf(msg, tt.input, tt.someone, tt.what, tt.when, someone, what, when)
		}
	}
}
