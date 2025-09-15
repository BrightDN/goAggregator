package rss

import (
	"net/http"
	"context"
	"io"
	"encoding/xml"
	"html"
	"time"

	"fmt"
)

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client {
	Timeout: 5 * time.Second,
 	}

	req.Header.Set("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { return nil, fmt.Errorf("bad status code: %s", resp.Status)}

	body, err := io.ReadAll(resp.Body)
	if err != nil { return nil, err}

	var rssfeed RSSFeed

	if err := xml.Unmarshal(body, &rssfeed); err != nil {
		return nil, err
	}

	ch := &rssfeed.Channel
	ch.Title = html.UnescapeString(ch.Title)
	ch.Description = html.UnescapeString(ch.Description)
	for i := range ch.Item {
		ch.Item[i].Title = html.UnescapeString(ch.Item[i].Title)
		ch.Item[i].Description = html.UnescapeString(ch.Item[i].Description)
	}

	return &rssfeed, nil
}