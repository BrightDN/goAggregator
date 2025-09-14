package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	AvailableCommands map[string]func(*config.State, Command) error
}

func (c *Commands) Run(s *config.State, cmd Command) error {
    if s == nil || s.Cfg == nil {
        return fmt.Errorf("invalid state")
    }

    if c.AvailableCommands == nil {
        return fmt.Errorf("no commands registered")
    }

    f, ok := c.AvailableCommands[cmd.Name]
    if !ok {
        return fmt.Errorf("unknown command: %s", cmd.Name)
    }

    if err := f(s, cmd); err != nil {
        return fmt.Errorf("command %q failed: %w", cmd.Name, err)
    }
	
	return nil
}

func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	c.AvailableCommands[name] = f
}