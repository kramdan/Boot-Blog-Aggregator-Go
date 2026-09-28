package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.RMUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Users table cleared\n")
	return nil
}
