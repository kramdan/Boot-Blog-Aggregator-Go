package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	var limit int32
	if len(cmd.args) != 1 {
		limit = 2
	} else {
		conv, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(conv)
	}
	posts, err := s.db.GetPostsForUser(context.Background(), limit)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("%v\n", post.Title)
		fmt.Printf("%v\n", post.PublishedAt.Time)
		fmt.Printf("%v\n", post.Description)
		fmt.Printf("\n")
	}

	return nil
}
