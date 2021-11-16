package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gritt/go-data-intensive/internal/core"
	"github.com/gritt/go-data-intensive/internal/details/kafka"
)

type Messenger struct {
	kafka *kafka.Client
}

func NewMessenger(kafka *kafka.Client) *Messenger {
	return &Messenger{kafka: kafka}
}

func (m *Messenger) Read(ctx context.Context, limit int) ([]core.Message, error) {
	data := m.kafka.ConsumeMessages()

	messages := []core.Message{}
	for _, item := range data {
		messages = append(messages, core.Message{
			Data:   item,
			UUID:   uuid.NewString(),
			Status: core.JOB_PENDING,
		})
	}

	return messages, nil
}

func (m *Messenger) Write(ctx context.Context, msg core.Message) error {
	return nil
}
