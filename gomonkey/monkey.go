package main

import (
	"fmt"
	"gomonkey/repl"
	"os"
	"os/user"
)

func main() {
	fmt.Println("Hello go")
	currentUser, _ := user.Current()
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", currentUser.Username)
	fmt.Println("Feel free to type in commands. Hit Control-D to exit.")
	repl.Start(os.Stdin, os.Stdout)
}
