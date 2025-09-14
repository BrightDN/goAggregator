package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
    "context"
)

func HandlerLogin(s *config.State, cmd Command) error {
    if len(cmd.Args) < 1 {
        return fmt.Errorf("login expects a username")
    }
    username := cmd.Args[0]

    ctx := context.Background()
	user, err := s.Db.GetUser(ctx, username)
    if err != nil { return err }

    if err := s.Cfg.SetUser(user.Name); err != nil {
        return fmt.Errorf("set user: %w", err)
    }

    fmt.Printf("user set to %s\n", username)
    return nil
}