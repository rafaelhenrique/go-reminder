package parser

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input   []string
	someone string
	what    string
	when    string
	err     error
}{
	{
		[]string{"remind", "me",
			"to", "drink", "water",
			"at", "3pm", "every", "day",
		},
		"me", "to drink water", "at 3pm every day", nil,
	},
	{
		[]string{"remind", "me",
			"to", "update", "the", "project", "status",
			"every", "Monday", "at", "9am",
		},
		"me", "to update the project status", "every Monday at 9am", nil,
	},
	{
		[]string{"remind", "whatever"},
		"", "", "", ErrSyntax,
	},
}

func TestParseArgs(t *testing.T) {
	for _, tt := range testCases {
		someone, what, when, err := ParseArgs(tt.input)

		assert.Equal(t, someone, tt.someone)
		assert.Equal(t, what, tt.what)
		assert.Equal(t, when, tt.when)
		assert.Equal(t, errors.Cause(err), errors.Cause(tt.err))
	}
}
