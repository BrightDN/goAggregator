package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
)

func HandlerFollow(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("missing a name and link argument")
	}

	link := cmd.Args[0]

	ctx := context.Background()

	feed, err := s.Db.GetFeedByURL(ctx, link)
	if err != nil {
		return err
	}

	feedfollow , err := s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}
	



	fmt.Println("User", feedfollow.Username)
	fmt.Println("Name:", feedfollow.Feedname)

    return nil
}