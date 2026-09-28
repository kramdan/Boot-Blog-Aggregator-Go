package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func handlerFetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var newFeed RSSFeed
	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		var blank RSSFeed
		return &blank, err
	}

	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		var blank RSSFeed
		return &blank, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		var blank RSSFeed
		return &blank, err
	}

	if err := xml.Unmarshal(body, &newFeed); err != nil {
		var blank RSSFeed
		return &blank, err
	}
	newFeed.Channel.Title = html.UnescapeString(newFeed.Channel.Title)
	newFeed.Channel.Description = html.UnescapeString((newFeed.Channel.Description))
	for _, item := range newFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
	return &newFeed, nil
}
