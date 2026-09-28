package main

import (
	config "Boot-Blog-Aggregator-Go/internal/config"
	"Boot-Blog-Aggregator-Go/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
