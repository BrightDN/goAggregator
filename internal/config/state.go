package config

import (
	"github.com/BrightDN/goAggregator/internal/database"
)

type State struct {
	Db *database.Queries
	Cfg *Config
}