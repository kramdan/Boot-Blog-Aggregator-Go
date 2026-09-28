package main

import (
	"Boot-Blog-Aggregator-Go/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments given, need username")
	}
	id := uuid.New()
	created_at := time.Now()
	updated_at := time.Now()
	name := cmd.args[0]
	check, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("Username %v already exists", check.Name)
	}
	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		Name:      name,
	}
	_, err = s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("User %v was created successfully\n", name)
	return nil
}
