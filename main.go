package main

import (
	"github.com/levicook/herbie/commands"
	_ "github.com/levicook/herbie/commands/server"
	"os"
)

func main() {
	command, err := commands.CommandFor(os.Args)

	if err != nil {
		panic(err)
	}

	command.Run()
}
