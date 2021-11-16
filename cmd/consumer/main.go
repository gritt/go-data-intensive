package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gritt/go-data-intensive/internal/core"
	"github.com/gritt/go-data-intensive/internal/service"
)

func main() {
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		initializeIndexingConsumers(ctx)
	}()

	wg.Wait()
	fmt.Println("main finished!")
}

func initializeIndexingConsumers(ctx context.Context) {
	numberOfWorkers := 5

	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		c := newIndexingConsumer(i)
		go c.Run(ctx, &wg)
	}

	wg.Wait()
}

func newIndexingConsumer(number int) Consumer {
	// TODO read envs...
	// TODO initialize dependencies...

	messenger := service.NewMessenger()
	searchIndex := core.NewSearchIndex()

	pipelineConfig := IndexingConfig{NumberOfWorkers: 10}
	pipeline := NewIndexingPipeline(
		pipelineConfig,
		messenger,
		searchIndex,
	)

	consumerConfig := Config{
		Name:     "indexing",
		UUID:     strconv.Itoa(number),
		Interval: 3 * time.Second,
	}
	return NewConsumer(consumerConfig, pipeline)
}
