package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	for i, com := range c.list {
		if cmd.name == i {
			err := com(s, cmd)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("Unknown or mistyped command\n")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}
