package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
	"fmt"
	"context"
)

func HandlerGetAllFollowings(s *config.State, cmd Command, user database.User) error {
	ctx := context.Background()

	feeds, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil { return err }

	for _, feed := range feeds {
		fmt.Println(feed.Feedname)
	}

    return nil
}