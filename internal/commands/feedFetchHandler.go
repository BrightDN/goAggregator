package commands

import (
	"github.com/BrightDN/goAggregator/internal/rss"
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
	"time"
	"context"
	"strings"
)

func HandlerFetchFeed(s *config.State, cmd Command) error {
    if len(cmd.Args) < 1 {
        return fmt.Errorf("Duration parameter is required, e.g. 1m, 1s...")
    }

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil { return err }

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	fmt.Printf("feeds will be fetched every %v\n", interval)
	for ; ; <-ticker.C {
        if err := scrapeFeeds(s); err != nil {
            fmt.Printf("Something went wrong: %v\n", err)
        }
	}
}

func scrapeFeeds(s *config.State) error {
	ctx := context.Background()
	feed, err := s.Db.GetNextFeedToFetch(ctx)
	
	if err != nil { return err }

	if err := s.Db.MarkFeedFetched(ctx, feed.ID); err != nil { return err }
	
	rssfeed, err := rss.FetchFeed(ctx, feed.Url)
    if err != nil { return err }

	for _, item := range rssfeed.Channel.Item {
		fmt.Printf("Title: %s\n", orDefault(item.Title, "missing title"))
	}

	fmt.Printf("\nEnd of feedpull for %s\n", rssfeed.Channel.Title)

    return nil
}

func orDefault(s, def string) string {
    if strings.TrimSpace(s) == "" {
        return def
    }
    return s
}