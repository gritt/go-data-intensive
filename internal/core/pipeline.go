package core

import (
	"context"
	"sync"
)

type Pipelined interface {
	Extract() ([]Job, error)
	Transform(jobs []Job) ([]WebPage, error)
	Load(webpages []WebPage) error
}

type Consumer interface {
	Run(withWaitGroup *sync.WaitGroup, ctx context.Context)
}
