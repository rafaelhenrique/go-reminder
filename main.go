package main

import (
	"fmt"
	"os"

	"github.com/rafaelhenrique/go-reminder/parser"
)

// "github.com/sqweek/dialog"

// func showMessage(msg string) {
// 	dialog.Message("%s", msg).Title("Reminder").Info()
// }

func main() {
	parsed := parser.ParseArgs(os.Args)
	fmt.Println(parsed)
}
