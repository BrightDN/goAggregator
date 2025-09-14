package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/database"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
)

func HandlerRegister(s *config.State, cmd Command) error {
    if len(cmd.Args) < 1 {
        return fmt.Errorf("login expects a username")
    }
    username := cmd.Args[0]

	ctx := context.Background()

	user, err := s.Db.CreateUser(ctx, database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	})

	if err != nil {
		return err
	}


    if err := s.Cfg.SetUser(user.Name); err != nil {
		return err
	}
    return nil
}