package main

import (
	_ "github.com/lib/pq"
	"github.com/BrightDN/goAggregator/internal/config"
	"github.com/BrightDN/goAggregator/internal/commands"
    "github.com/BrightDN/goAggregator/internal/database"
	"fmt"
	"os"
    "database/sql"
)


func main() {
    cfg, err := config.Read()
    if err != nil {
        fmt.Fprintf(os.Stderr, "read config: %v\n", err)
        os.Exit(1)
    }

    db, err := sql.Open("postgres", cfg.DBURL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Database connection: %w\n", err)
    }

    defer db.Close()

    dbQueries := database.New(db)

    state := config.State{
        Db: dbQueries,
        Cfg: &cfg,
    }
    cmds := commands.Commands{
        AvailableCommands: make(map[string]func(*config.State, commands.Command) error),
    }
    cmds.Register("login", commands.HandlerLogin)
    cmds.Register("register", commands.HandlerRegister)
    cmds.Register("reset", commands.HandlerReset)

    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "not enough arguments")
        os.Exit(1)
    }

    cmd := commands.Command{
        Name: os.Args[1],
        Args: os.Args[2:],
    }

    if err := cmds.Run(&state, cmd); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
