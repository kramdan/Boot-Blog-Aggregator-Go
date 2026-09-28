package main

import (
	"context"
	"fmt"
)

func handlerAgg(state *state, cmd command) error {
	rssFeed, err := handlerFetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", rssFeed)
	return nil
}
