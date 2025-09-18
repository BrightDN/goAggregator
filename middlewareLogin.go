package main

import (
	"context"
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/commands"
	"github.com/BrightDN/goAggregator/internal/database"
)
func middlewareLoggedIn(
    handler func(s *config.State, cmd commands.Command, user database.User) error,
) func(*config.State, commands.Command) error {
    return func(s *config.State, cmd commands.Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}
        return handler(s, cmd, user)
    }
}