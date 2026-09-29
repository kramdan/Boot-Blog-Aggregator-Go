package main

import (
	"Boot-Blog-Aggregator-Go/internal/database"
	"context"
	"fmt"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Need URL of feed to unfollow")
	}
	feed, err := s.db.GetFeedURL(context.Background(), cmd.args[0])
	params := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	_, err = s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %v unfollowed", feed.Name)
	return nil
}
