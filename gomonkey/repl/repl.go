package repl

import (
	"bufio"
	"fmt"
	"gomonkey/lexer"
	"gomonkey/token"
	"io"
)

const PROMPT = "🐒 > "

// Start this procedure starts REPL
func Start(in io.Reader, _ io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		} else {
			line := scanner.Text()
			l := lexer.New(line)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		}
	}
}
