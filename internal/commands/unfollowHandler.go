package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
	"fmt"
    "context"
)

func Handlerunfollow(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("expects an URL to be given as parameter")
	}
    ctx := context.Background() 
	feed, err := s.Db.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil { return err }

    if err := s.Db.DeleteByUrl(ctx, database.DeleteByUrlParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil { return err }
    return nil
}