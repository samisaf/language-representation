package main

import (
	"fmt"
	"gomonkey/lexer"
	"os/user"
)

func main() {
	fmt.Println("Hello go")
	currentUser, _ := user.Current()
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", currentUser.Username)
	fmt.Println("Feel free to type in commands. Hit Control-C or Control-D to exit.")
	testLexer()
	// repl.Start(os.Stdin, os.Stdout)
}

func testLexer() {
	st1 := "let a = 5;"
	l1 := lexer.New(st1)
	for tok := l1.NextToken(); tok.Type != "EOF"; tok = l1.NextToken() {
		fmt.Println(tok)
	}
}
