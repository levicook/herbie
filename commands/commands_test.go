package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type commandA struct{}
type commandB struct{}

func (commandA) Name() string { return "commandA" }
func (commandB) Name() string { return "commandB" }

func (commandA) Run() {}
func (commandB) Run() {}

func Test(t *testing.T) {
	Register(commandA{})
	Register(commandB{})

	assert.Panics(t, func() { Register(commandA{}) })
	assert.Panics(t, func() { Register(commandB{}) })

	var (
		cmd Command
		err error
	)

	cmd, err = CommandFor([]string{"herbie", "commandA"})
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name(), "commandA")

	cmd, err = CommandFor([]string{"herbie", "commandB"})
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name(), "commandB")

	cmd, err = CommandFor([]string{"herbie", "commandC"})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), `command "commandC" not found`)
}
