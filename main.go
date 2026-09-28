package main

import (
	config "Boot-Blog-Aggregator-Go/internal/config"
	"Boot-Blog-Aggregator-Go/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	dbURL := conf.DBURL
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	if err != nil {
		fmt.Printf("%v", err)
	}
	mainState := state{
		db:  dbQueries,
		cfg: &conf,
	}
	commandsList := commands{
		list: make(map[string]func(*state, command) error),
	}
	commandsList.register("login", handlerLogin)
	commandsList.register("register", handlerRegister)
	commandsList.register("reset", handlerReset)
	commandsList.register("users", handlerUsers)
	commandsList.register("agg", handlerAgg)
	commandsList.register("addfeed", handlerAddFeed)
	commandsList.register("feeds", handlerFeeds)
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Not enough arguments\n")
		os.Exit(1)
	}
	cmd := args[1]
	cmdArgs := args[2:]
	err = commandsList.run(&mainState, command{name: cmd, args: cmdArgs})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
