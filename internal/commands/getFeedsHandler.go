package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
	"context"
)

func HandlerGetFeeds(s *config.State, cmd Command) error {
	ctx := context.Background()
	feeds, err := s.Db.GetAllFeeds(ctx)
	if err != nil { return err }

	for _, feed := range feeds {
		fmt.Printf("Feedname: %s\nFeed url: %s\nFeed creator: %s\n\n", feed.Name, feed.Url, feed.Username)
	}

    return nil
}