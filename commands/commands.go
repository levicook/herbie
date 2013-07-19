package commands

import (
	"fmt"
)

type Command interface {
	Name() string
	Run()
}

var registry = make(map[string]Command)

func CommandFor(args []string) (command Command, err error) {
	var (
		name string
		ok   bool
	)

	if len(args) > 1 {
		name = args[1]
	}

	if command, ok = registry[name]; !ok {
		err = fmt.Errorf("command %q not found", name)
	}

	return
}

func Register(command Command) {
	if _, ok := registry[command.Name()]; ok {
		panic(fmt.Errorf("command %q is already in use", command.Name()))
	}

	registry[command.Name()] = command
}
