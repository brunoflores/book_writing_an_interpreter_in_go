package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
	"strings"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		reader := strings.NewReader(line)
		l := lexer.New(reader, "-")
		for {
			tok := l.NextToken()
			fmt.Printf("%+v\n", tok)

			if token.Is(tok, token.Symbol{Id: token.EOF}) {
				break
			}
		}
	}
}
