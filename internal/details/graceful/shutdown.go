package graceful

import (
	"context"
	"os"
	"os/signal"
)

func ShutdownWith(ctx context.Context, exit chan bool) {
	ctx, cancel := context.WithCancel(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig

		cancel()

		exit <- true
	}()
}
