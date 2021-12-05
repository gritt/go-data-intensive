package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gritt/go-data-intensive/internal/core"
)

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
	return p.messenger.Read(ctx, 2)
}

func (p *IndexingPipeline) Transform(ctx context.Context, msgs []core.Message) ([]core.WebPage, error) {
	pendingJobs := make(chan core.Message, len(msgs))
	completedJobs := make(chan core.Message, len(msgs))

	webpages := []core.WebPage{}

	// non blocking - concurrency
	for i := 0; i < p.config.NumberOfWorkers; i++ {
		go func() {

			// read jobs until pending channel is empty
			for message := range pendingJobs {

				fmt.Printf("→: %s \n", message.UUID)

				webpage, err := p.indexer.Process(ctx, message.Data)
				if err != nil {
					// TODO err handling
					fmt.Println(err)
				} else {
					webpages = append(webpages, webpage)
					time.Sleep(time.Second * 1)
				}

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
		fmt.Printf("✔: %s \n", msg.UUID)
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
