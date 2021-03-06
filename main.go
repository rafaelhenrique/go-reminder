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
	someone, what, when, err := parser.ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Someone: %+v\n", someone)
	fmt.Printf("What: %+v\n", what)
	fmt.Printf("When: %+v\n", when)
}
