package main

import (
	config "Boot-Blog-Aggregator-Go/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
	}
	user := "daniel"
	cfg.SetUser(user)
	updatedConfig, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", updatedConfig)
}
