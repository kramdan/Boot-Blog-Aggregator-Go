package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	list, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, name := range list {
		currentUser := s.cfg.CurrentUserName
		if name == currentUser {
			fmt.Printf("* %v (current)\n", name)
			continue
		}
		fmt.Printf("* %v\n", name)
	}
	return nil
}
