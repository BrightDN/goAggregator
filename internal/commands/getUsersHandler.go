package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
	"context"
)

func HandlerGetUsers(s *config.State, cmd Command) error {
	ctx := context.Background()
	users, err := s.Db.GetUsers(ctx)
	if err != nil { return err }

	for _, user := range users {
		currentUserTxt := ""
		if user.Name == s.Cfg.CurrentUserName {
			currentUserTxt = "(current)"
		}
		fmt.Printf("* %s %s\n", user.Name, currentUserTxt)
	}

    return nil
}