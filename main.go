package main

import (
	"fmt"
	"github.com/negarciacamilo/go_interpreter/repl"
	"os"
	user "os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Type in some commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
