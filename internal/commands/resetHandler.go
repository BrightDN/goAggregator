package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
    "context"
)

func HandlerReset(s *config.State, cmd Command) error {
    ctx := context.Background() 
    if err := s.Db.ResetUsers(ctx); err != nil { return err }
    fmt.Println("Table succesfully emptied")
    return nil
}