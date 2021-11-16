package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gritt/go-data-intensive/internal/core"
	"github.com/gritt/go-data-intensive/internal/details/graceful"
)

type (
	Pipelined interface {
		Extract(ctx context.Context) ([]core.Message, error)
		Transform(ctx context.Context, msgs []core.Message) ([]core.WebPage, error)
		Load(ctx context.Context, webpages []core.WebPage) error
	}

	Config struct {
		Name     string
		UUID     string
		Interval time.Duration
	}

	Consumer struct {
		config   Config
		pipeline Pipelined
	}
)

func NewConsumer(cfg Config, pipeline Pipelined) Consumer {
	return Consumer{
		config:   cfg,
		pipeline: pipeline,
	}
}

func (c Consumer) Run(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Printf("Start: consumer_%s_%s \n", c.config.Name, c.config.UUID)

	defer func() {
		wg.Done()
		fmt.Printf("Exit: consumer_%s_%s \n", c.config.Name, c.config.UUID)
	}()

	exit := make(chan bool)
	go func() {
		graceful.ShutdownWith(ctx, exit)
	}()

	for {
		msgs, err := c.pipeline.Extract(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}

		webpages, err := c.pipeline.Transform(ctx, msgs)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = c.pipeline.Load(ctx, webpages)
		if err != nil {
			fmt.Println(err.Error())
		}

		select {
		case <-exit:
			return
		case <-time.After(c.config.Interval):
			continue
		}
	}
}
