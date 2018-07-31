package main

import (
	"strings"

	"github.com/sqweek/dialog"
)

func showMessage(msg string) {
	dialog.Message("%s", msg).Title("Reminder").Info()
}

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

func main() {

}
