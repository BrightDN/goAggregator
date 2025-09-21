package commands

import (
	"github.com/BrightDN/goAggregator/internal/rss"
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

		pubAt, err := parseAny(item.PubDate)
		if err != nil { fmt.Errorf("Error formatting date: %w", err) }
		pubAt = pubAt.UTC()


		if err := s.Db.CreatePost(ctx, database.CreatePostParams{
			ID: uuid.New(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: pubAt,
			FeedID: feed.ID,
		}); err != nil {
    		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code != "23505" { 
				return err
			}
		}
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

var layouts = []string{
    time.RFC3339, time.RFC3339Nano,
    time.RFC1123Z, time.RFC1123,
    time.RFC822Z, time.RFC822,
    "Mon, 02 Jan 2006 15:04:05 MST",
}

func parseAny(ts string) (time.Time, error) {
    for _, l := range layouts {
        if t, err := time.Parse(l, ts); err == nil {
            return t, nil
        }
    }
    return time.Time{}, fmt.Errorf("unrecognized time: %q", ts)
}