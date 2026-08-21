package main

import (
	"log"

	"github.com/saurav11sarkar/ticket/internal/config"
	"github.com/saurav11sarkar/ticket/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := config.ConnectDatabase(cfg)

	if err := server.Start(db, cfg); err != nil {
		log.Fatal(err)
	}
}
