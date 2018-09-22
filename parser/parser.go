package parser

import (
	"github.com/pkg/errors"

	"strings"
)

var (
	// ErrSyntax syntax error on ParseArgs
	ErrSyntax = errors.New("Syntax error")
)

// ParseArgs receive an slice of strings and return someone, what and when
func ParseArgs(args []string) (someone, what, when string, err error) {
	for _, argument := range args[1:] {
		if argument == "me" {
			someone = argument
		} else if argument == "at" || argument == "every" || when != "" {
			when += argument + " "
		} else {
			what += argument + " "
		}
	}
	someone = strings.Trim(someone, " ")
	what = strings.Trim(what, " ")
	when = strings.Trim(when, " ")

	if someone == "" || what == "" || when == "" {
		err = errors.Wrapf(
			ErrSyntax,
			"package=parser, function=ParseArgs, error=ErrSyntax, args=%#v",
			args,
		)

		return "", "", "", err
	}

	return
}
