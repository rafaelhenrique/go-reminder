package parser

import "strings"

// ParseArgs receive an slice of strings and return someone, what and when
func ParseArgs(args []string) (someone, what, when string) {
	for _, argument := range args[1:] {
		if argument == "me" {
			someone = argument
		} else if argument == "at" || argument == "every" || when != "" {
			when += argument + " "
		} else {
			what += argument + " "
		}
	}
	return strings.Trim(someone, " "),
		strings.Trim(what, " "),
		strings.Trim(when, " ")
}
