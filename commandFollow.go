package main

import (
	"Boot-Blog-Aggregator-Go/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Need url for feed")
	}
	feed, err := s.db.GetFeedURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %v followed by User: %v\n", follow.FeedName, follow.UserName)
	return nil
}
