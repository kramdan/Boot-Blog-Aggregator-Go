package main

import (
	"Boot-Blog-Aggregator-Go/internal/database"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Need duration string of #s, or #m")
	}
	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collection feeds every %v", duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err = scrapeFeeds(state)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.NextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	markedFeed, err := s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}
	fetchedFeed, err := handlerFetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return err
	}
	for _, item := range fetchedFeed.Channel.Item {
		descr := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}
		date, err := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700", item.PubDate)
		if err != nil {
		}
		pubdate := sql.NullTime{
			Time:  date,
			Valid: true,
		}
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: descr,
			PublishedAt: pubdate,
			FeedID:      markedFeed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), params)

	}
	return nil
}
