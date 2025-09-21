package commands

import (
	"context"
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
)
func middlewareLoggedIn(
    handler func(s *config.State, cmd Command, user database.User) error,
) func(*config.State, Command) error {
    return func(s *config.State, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}
        return handler(s, cmd, user)
    }
}