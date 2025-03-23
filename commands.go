package main

import (
	"fmt"

	"github.com/breeze/blogagg/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: login <username>")
	}
	fmt.Println("User set to", cmd.args[0])
	return s.config.SetUser(cmd.args[0])
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command %q", cmd.name)
	}
	return f(s, cmd)
}
