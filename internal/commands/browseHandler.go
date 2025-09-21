package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
	"fmt"
	"context"
	"strconv"
)

func HandlerBrowse(s *config.State, cmd Command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) >= 1 {
		n64, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil { return fmt.Errorf("Cannot parse the limit: %w", err) }
		limit = int32(n64)
	}

	ctx := context.Background()

	posts, err := s.Db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: limit,
	})
	if err != nil { return err }

	for _, post := range posts {
		fmt.Printf("Title: %s\nLink: %s\n", post.Title, post.Url)
	}

    return nil
}