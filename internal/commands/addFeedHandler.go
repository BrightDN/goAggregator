package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *config.State, cmd Command, user database.User) error {
	name := ""
	link := ""
	switch len(cmd.Args) {
	case 0:
		return fmt.Errorf("missing a name and link argument")
	case 1:
		return fmt.Errorf("missing a link argument")
	default:
		name = cmd.Args[0]
		link = cmd.Args[1]
	}

	ctx := context.Background()

	feed , err := s.Db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: link,
		UserID: user.ID,
	})
	
	if err != nil {
		return err
	}

	_ , ffErr := s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if ffErr != nil {
		return ffErr
	}

	fmt.Println("ID:", feed.ID)
	fmt.Println("CreatedAt:", feed.CreatedAt)
	fmt.Println("UpdatedAt:", feed.UpdatedAt)
	fmt.Println("Name:", feed.Name)
	fmt.Println("URL:", feed.Url)
	fmt.Println("UserID:", feed.UserID)

    return nil
}