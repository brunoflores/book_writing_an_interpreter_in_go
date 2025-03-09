package main

import (
	"interpreter/lexer"
	"interpreter/repl"
	"interpreter/token"

	"fmt"
	"os"
	"os/user"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		fname := args[0]
		file, err := os.Open(fname)
		if err != nil {
			panic(err)
		}
		l := lexer.New(file, fname)
		for {
			tok := l.NextToken()
			fmt.Printf("%+v\n", tok)

			if token.Is(tok, token.Symbol{Id: token.EOF}) {
				break
			}
		}

		return
	}

	// Running the repl
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
