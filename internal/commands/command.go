package commands

import (
	"github.com/BrightDN/goAggregator/internal/config"
	"fmt"
)

type Command struct {
	Name string
	Args []string

    Description string
    ArgHelp     map[string]string 
    Run         func(*config.State, Command) error           
}

type Commands struct {
    registry map[string]func(*config.State, Command) error
    metas    map[string]Command
}

func New() *Commands {
    return &Commands{
        registry: make(map[string]func(*config.State, Command) error),
        metas:    make(map[string]Command),
    }
}

func (c *Commands) Run(s *config.State, cmd Command) error {
    if s == nil || s.Cfg == nil {
        return fmt.Errorf("invalid state")
    }

    if c.registry == nil {
        return fmt.Errorf("no commands registered")
    }

    f, ok := c.registry[cmd.Name]
    if !ok {
        return fmt.Errorf("unknown command: %s", cmd.Name)
    }

    if err := f(s, cmd); err != nil {
        return fmt.Errorf("command %q failed: %w", cmd.Name, err)
    }
	
	return nil
}

func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	c.registry[name] = f
}

func (c *Commands) RegisterAll() {
    for _, cmd := range AllCommands {
        c.registry[cmd.Name] = cmd.Run
        c.metas[cmd.Name] = cmd
    }
}

var AllCommands = []Command{
    {
        Name:        "reset",
        Description: "Resets the database",
        ArgHelp: map[string]string{},
        Run: HandlerReset,
    },
    {
        Name:        "login",
        Description: "Log in with your username",
        ArgHelp: map[string]string{
            "username": "The username you wish to log in with",
        },
        Run: HandlerLogin,
    },
    {
        Name:        "register",
        Description: "Register a username",
        ArgHelp: map[string]string{
            "username": "The username you wish to register as",
        },
        Run: HandlerRegister,
    },
    {
        Name:        "users",
        Description: "Return a list of all the users known to the system",
        ArgHelp: map[string]string{},
        Run: HandlerGetUsers,
    },
    {
        Name:        "agg",
        Description: "Fetch feeds with a set interval",
        ArgHelp: map[string]string{
            "interval": "Interval between fetch requests (e.g. 1m, 5s, ...)",
        },
        Run: HandlerFetchFeed,
    },
    {
        Name:        "feeds",
        Description: "Get a list of all the feeds currently in the system",
        ArgHelp: map[string]string{},
        Run: HandlerGetFeeds,
    },
    {
        Name:        "addfeed",
        Description: "Add a feed to the list of feeds",
        ArgHelp: map[string]string{
            "name": "The name of the feed",
            "url": "The link to the RSSfeed",
        },
        Run: middlewareLoggedIn(HandlerAddFeed),
    },
    {
        Name:        "following",
        Description: "Get a list of all the feeds followed by the current user",
        ArgHelp: map[string]string{},
        Run: middlewareLoggedIn(HandlerGetAllFollowings),
    },
    {
        Name:        "follow",
        Description: "Start following a feed",
        ArgHelp: map[string]string{
            "url": "The link to the RSSfeed you wish to follow",
        },
        Run: middlewareLoggedIn(HandlerFollow),
    },
    
    {
        Name:        "unfollow",
        Description: "stop following a feed",
        ArgHelp: map[string]string{
            "url": "The link to the RSSfeed you wish to stop following",
        },
        Run: middlewareLoggedIn(Handlerunfollow),
    },
    {
        Name:        "browse",
        Description: "browse through your feeds",
        ArgHelp: map[string]string{
            "limit": "[OPTIONAL] Add a limit of feeds to see. Defaults to 2 if missing",
        },
        Run: middlewareLoggedIn(HandlerBrowse),
    },
    
}
