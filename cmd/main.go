package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/gritt/go-data-intensive/internal/services/consumer"
)

func main() {
	// TODO init ctx
	ctx := context.Background()

	// TODO init shared dependencies
	// repositories, services

	// TODO create other workers

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		initializeCrawlerWorkers(ctx)
	}()

	wg.Wait()
	fmt.Println("main finished!")
}

func initializeCrawlerWorkers(ctx context.Context) {
	// TODO read number of workers from env
	numberOfWorkers := 5

	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		c := initializeCrawlerConsumer(i)
		go c.Run(&wg, ctx)
	}

	wg.Wait()
}

func initializeCrawlerConsumer(number int) consumer.CrawlerConsumer {
	// TODO add interval to config
	// TODO read number of workers from env

	name := fmt.Sprintf("CRAWLER_CONSUMER_%d", number)

	cfg := consumer.Config{
		UUID:            name,
		NumberOfWorkers: 50,
	}

	return consumer.NewCrawlerConsumer(cfg)
}
