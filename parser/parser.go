package parser

import "strings"

// parseArgs receive an slice of strings and return an reminder phrase
func parseArgs(args []string) (reminder string) {
	var someone, what, when string

	for _, argument := range args[1:] {
		if argument == "me" {
			someone = argument
		} else if argument == "at" || argument == "every" || when != "" {
			when += argument + " "
		} else {
			what += argument + " "
		}
	}
	reminder = strings.Trim(someone, " ") +
		"|" + strings.Trim(what, " ") +
		"|" + strings.Trim(when, " ")
	return
}
