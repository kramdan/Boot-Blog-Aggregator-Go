package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments given, need username")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("User not found")
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("The user %s has been set\n", cmd.args[0])
	return nil
}
