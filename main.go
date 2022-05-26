package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	var name string
	if user.Name == "" {
		name = user.Username
	} else {
		name = user.Name
	}

	fmt.Printf("Hello %s! This is the Monkey programming language\n", name)
	fmt.Println("Feel free to type in commands")
	if len(os.Args) > 1 {
		t := os.Args[1]
		if t == "interpreter" || t == "-i" {
			repl.StartInterpreter(os.Stdin, os.Stdout)
		} else if t == "compiler" || t == "-c" {
			repl.Start(os.Stdin, os.Stdout)
		}
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
