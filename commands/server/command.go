package server

import (
	"github.com/levicook/herbie/commands"
	"log"
)

func init() { commands.Register(server{}) }

type server struct{}

func (server) Name() string { return "server" }

func (cmd server) Run() {
	log.Println(cmd.Name())
}
