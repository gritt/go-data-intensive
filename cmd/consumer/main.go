package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gritt/go-data-intensive/internal/core"
	"github.com/gritt/go-data-intensive/internal/details/kafka"
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
	fmt.Println("main exited gracefully!")
}

func initializeIndexingConsumers(ctx context.Context) {
	numberOfWorkers := 1

	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		c := newIndexingConsumer(i)
		go c.Run(ctx, &wg)
	}

	wg.Wait()
}

func newIndexingConsumer(number int) Consumer {
	kafka := newKafkaClient()

	messenger := service.NewMessenger(kafka)

	searchIndex := core.NewSearchIndex()

	pipeline := newIndexingPipeline(messenger, searchIndex)

	consumerConfig := Config{
		Name:     "indexing",
		UUID:     strconv.Itoa(number),
		Interval: 10 * time.Millisecond,
	}
	return NewConsumer(consumerConfig, pipeline)
}

func newKafkaClient() *kafka.Client {
	// TODO read from env
	cfg := kafka.Config{
		MaxRetries: 3,
		Host:       "localhost:9092",
		Topic:      "test",
		Partition:  int32(0),
	}

	return kafka.NewClient(cfg)
}

func newIndexingPipeline(m core.Messenger, i core.Indexer) *IndexingPipeline {
	// TODO read from env
	cfg := IndexingConfig{
		NumberOfWorkers: 10,
	}

	return NewIndexingPipeline(cfg, m, i)
}
