package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/gritt/go-data-intensive/internal/core"
)

type Messenger struct {
}

func NewMessenger() *Messenger {
	return &Messenger{}
}

func (m *Messenger) Read(ctx context.Context, limit int) ([]core.Message, error) {
	fakeMsgs := []core.Message{}

	for j := 0; j <= limit; j++ {
		fakeMsgs = append(fakeMsgs, core.Message{
			Data:   fakeString(100),
			UUID:   fakeString(10),
			Status: core.JOB_PENDING,
		})
	}

	return fakeMsgs, nil
}

func (m *Messenger) Write(ctx context.Context, msg core.Message) error {
	return nil
}

func (m *Messenger) Delete(ctx context.Context, msg core.Message) error {
	return nil
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
