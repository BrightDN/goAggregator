package commands

import (
	"github.com/BrightDN/goAggregator/internal/rss"
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
	"context"
)

func HandlerFetchFeed(s *config.State, cmd Command) error {
    // if len(cmd.Args) < 1 {
    //     return fmt.Errorf("agg expects an url")
    // }
    // url := cmd.Args[0]

	url := "https://www.wagslane.dev/index.xml"
    ctx := context.Background()
	rssfeed, err := rss.FetchFeed(ctx, url)
    if err != nil { return err }

	fmt.Printf("%+v\n", *rssfeed) 

    return nil
}

