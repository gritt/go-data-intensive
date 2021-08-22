package consumer

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gritt/go-data-intensive/internal/core"
	"github.com/gritt/go-data-intensive/internal/services/graceful"
)

type Config struct {
	UUID            string
	NumberOfWorkers int
}

type CrawlerConsumer struct {
	Config
}

func NewCrawlerConsumer(cfg Config) CrawlerConsumer {
	return CrawlerConsumer{
		Config: cfg,
	}
}

func (s CrawlerConsumer) Run(wg *sync.WaitGroup, ctx context.Context) {
	fmt.Println(fmt.Sprintf("Starting CrawlerConsumer %s", s.UUID))

	defer func() {
		wg.Done()
		fmt.Println(fmt.Sprintf("Finish CrawlerConsumer %s", s.UUID))
	}()

	exit := make(chan bool)
	go func() {
		graceful.ShutdownWith(ctx, exit)
	}()

	for {
		jobs, err := s.Extract()
		if err != nil {
			fmt.Println(err.Error())
		}

		webpages, err := s.Transform(jobs)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = s.Load(webpages)
		if err != nil {
			fmt.Println(err.Error())
		}

		// TODO read wait interval from config
		select {
		case <-exit:
			return
		case <-time.After(3 * time.Second):
			continue
		}
	}
}

func (s CrawlerConsumer) Extract() ([]core.Job, error) {
	// TODO read/add number of jobs to collect per execution to config

	// lê um stream com lista de links, ou inicializa com entrypoint link (hacker news)
	// retorna html das páginas web
	// cria jobs com essas paginas

	fakeJobs := collectFakeJobs(100)

	return fakeJobs, nil
}

func (s CrawlerConsumer) Transform(jobs []core.Job) ([]core.WebPage, error) {
	pending := make(chan core.Job, len(jobs))
	completed := make(chan core.Job, len(jobs))

	// process with concurrency
	for i := 1; i < s.NumberOfWorkers; i++ {
		go func() {
			for pendingJob := range pending {

				// gerar webpage a partir dos dados string de cada job

				fmt.Printf("process job: %s \n", pendingJob.UUID)

				time.Sleep(time.Second * 1)

				completed <- pendingJob
			}
		}()
	}

	// enqueue jobs
	for _, job := range jobs {
		pending <- job
	}
	close(pending)

	// get results
	for r := 1; r <= len(jobs); r++ {
		job := <-completed
		fmt.Printf("job %s  ✔ \n", job.UUID)
	}
	close(completed)

	return []core.WebPage{}, nil
}

func (s CrawlerConsumer) Load(webpages []core.WebPage) error {

	// indexa ELK

	// adiciona todos os links de todas as páginas encontradas no stream de links

	fmt.Println(webpages)
	return nil
}

func collectFakeJobs(amount int) []core.Job {
	jobs := []core.Job{}

	for j := 0; j <= amount; j++ {
		jobs = append(jobs, core.Job{
			Data:   fakeString(100),
			UUID:   fakeString(10),
			Status: core.JOB_PENDING,
		})
	}

	return jobs
}

func fakeString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
