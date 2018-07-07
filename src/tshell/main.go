package main

import (
	"fmt"
	"os/user"
	"tshell/editor"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s! You can translate input languages.\n",
		user.Username)
	editor.Start()
}
