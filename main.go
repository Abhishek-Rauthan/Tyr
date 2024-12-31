package main

import (
	"fmt"
	"os"
	"os/user"
	"tyr/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Print(repl.TYR_FACE)
	fmt.Printf("Hello %s! This is the tyr programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
