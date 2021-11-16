package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type message struct {
	Data string `json:"data"`
	UUID string `json:"uuid"`
}

type Config struct {
	MaxRetries int
	Host       string
	Topic      string
	Partition  int32
}

type Client struct {
	config     Config
	producer   sarama.SyncProducer
	consumer   sarama.Consumer
	lastOffset int64
}

func NewClient(cfg Config) *Client {
	hosts := []string{cfg.Host}
	saramaConfig := sarama.NewConfig()

	// producer
	saramaConfig.Producer.Retry.Max = cfg.MaxRetries
	saramaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(hosts, saramaConfig)
	if err != nil {
		fmt.Println("NewAsyncProducer err:")
		fmt.Println(err)
	}

	// consumer
	saramaConfig.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(hosts, saramaConfig)
	if err != nil {
		fmt.Println("NewConsumer err:")
		fmt.Println(err)
	}

	return &Client{
		config:     cfg,
		producer:   producer,
		consumer:   consumer,
		lastOffset: sarama.OffsetOldest,
	}
}

func (c *Client) ConsumeMessages() []string {
	consumerPartition, err := c.consumer.ConsumePartition(
		c.config.Topic,
		c.config.Partition,
		c.lastOffset,
	)
	if err != nil {
		// TODO err handling
		fmt.Println(err)
	}
	defer consumerPartition.Close()

	data := []string{}
	for {
		select {
		case err := <-consumerPartition.Errors():
			fmt.Println("Error", err.Error())

		case msg := <-consumerPartition.Messages():
			if msg.Offset != c.lastOffset {
				data = append(data, string(msg.Value))
				c.lastOffset = msg.Offset
			}
		case <-time.After(10 * time.Millisecond):
			return data
		}
	}
}
