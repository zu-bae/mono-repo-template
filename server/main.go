package main

import (
	"log"

	"github.com/zu-bae/v-connect/server/config"
	"github.com/zu-bae/v-connect/server/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := router.New(cfg)
	log.Printf("env=%s listening on %s", cfg.Env, cfg.Addr())
	if err := r.Run(cfg.Addr()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
