package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gritt/go-data-intensive/internal/core"
)

// TODO error handling
// TODO channel for failed jobs
// TODO retry logic
// TODO fallback / dead letter queue / no ack
// TODO circuit breaker

type (
	IndexingConfig struct {
		NumberOfWorkers int
	}

	IndexingPipeline struct {
		config    IndexingConfig
		messenger core.Messenger
		indexer   core.Indexer
	}
)

func NewIndexingPipeline(cfg IndexingConfig, m core.Messenger, i core.Indexer) *IndexingPipeline {
	return &IndexingPipeline{
		config:    cfg,
		messenger: m,
		indexer:   i,
	}
}

func (p *IndexingPipeline) Extract(ctx context.Context) ([]core.Message, error) {
	return p.messenger.Read(ctx, 50)
}

func (p *IndexingPipeline) Transform(ctx context.Context, msgs []core.Message) ([]core.WebPage, error) {
	pendingJobs := make(chan core.Message, len(msgs))
	completedJobs := make(chan core.Message, len(msgs))

	webpages := []core.WebPage{}

	// non blocking - concurrency
	for i := 1; i < p.config.NumberOfWorkers; i++ {
		go func() {
			// read jobs until pending channel is empty
			for message := range pendingJobs {

				// message is processing
				fmt.Printf("processing: %s \n", message.UUID)

				webpage, err := p.indexer.Process(ctx, message.Data)
				if err != nil {
					// TODO err handling
					fmt.Println(err)
				} else {
					webpages = append(webpages, webpage)
					time.Sleep(time.Second * 1)
				}

				// message was processed
				completedJobs <- message
			}
		}()
	}

	// blocking / enqueue jobs
	for _, msg := range msgs {
		pendingJobs <- msg
	}
	close(pendingJobs)

	// blocking / log finished jobs
	for r := 1; r <= len(msgs); r++ {
		msg := <-completedJobs
		fmt.Printf("job %s  ✔ \n", msg.UUID)
	}
	close(completedJobs)

	return webpages, nil
}

func (p *IndexingPipeline) Load(ctx context.Context, webpages []core.WebPage) error {
	// TODO concurrency
	for _, webpage := range webpages {

		if err := p.indexer.Store(ctx, webpage); err != nil {
			// TODO err handling
			fmt.Println(err)
		}
	}
	return nil
}
