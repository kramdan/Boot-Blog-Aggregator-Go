package main

import (
	"Boot-Blog-Aggregator-Go/internal/database"
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feed := range following {
		fmt.Printf("%v\n", feed.FeedName)
	}
	return nil
}
