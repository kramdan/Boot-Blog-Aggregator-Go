package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	list, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range list {
		creator, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %v, URL: %v, Creator: %v\n", feed.Name, feed.Url, creator)
	}
	return nil
}
