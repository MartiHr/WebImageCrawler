package crawler

import (
	"Homework2/internal/storage"
	"context"
	"fmt"
)

type Engine struct {
	Workers        int
	MaxGoroutines  int
	FollowExternal bool

	jobs    chan string
	visited *URLSet
	sem     chan struct{}

	DB *storage.ImageRepository
}

func NewEngine(workers, maxGoroutines int) *Engine {
	return &Engine{
		Workers:       workers,
		MaxGoroutines: maxGoroutines,
		jobs:          make(chan string, 1000),
		visited:       NewURLSet(),
		sem:           make(chan struct{}, maxGoroutines),
	}
}

func (e *Engine) Start(ctx context.Context, startURLs []string) {
	for i := 0; i < e.Workers; i++ {
		go e.worker(ctx, i)
	}

	for _, url := range startURLs {
		e.jobs <- url
	}

	<-ctx.Done()
	close(e.jobs)
	fmt.Println("Crawler stopped")
}
