package main

import (
	"Homework2/internal/config"
	"Homework2/internal/crawler"
	"Homework2/internal/storage"
	"Homework2/internal/web"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {

	cfg := config.Load()

	if len(cfg.StartURLs) == 0 {
		fmt.Println("Usage: crawler [options] <url1> <url2> ...")
		os.Exit(1)
	}

	// Context ONLY for crawling
	crawlCtx, cancel := context.WithTimeout(
		context.Background(),
		cfg.Timeout,
	)
	defer cancel()

	// DB
	db, err := storage.Open(cfg.MySQLDSN)
	if err != nil {
		log.Fatal(err)
	}

	repo := storage.NewImageRepository(db)

	// Crawler
	engine := crawler.NewEngine(cfg.Workers, cfg.MaxGoroutines)
	engine.DB = repo
	engine.FollowExternal = cfg.FollowExternal

	go engine.Start(crawlCtx, cfg.StartURLs)

	// Web UI (no timeout)
	web.Init(repo)

	log.Println("Web UI at http://localhost:8080")
	log.Fatal(web.Start(":8080"))
}

// go run . -workers=3 -max-goroutines=5 https://example.com https://golang.org
// go run . -workers=10 -max-goroutines=50 -timeout=1m -external=false  https://example.com https://golang.org

// For UI testing
// URL contains: golang.org
// Filename contains: gopher
// Alt text: logo
// Title: Go
// Format: png
// Min width: 100
// Max width: 300
// Min height: 100
// Max height: 300
