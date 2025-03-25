package govnr

import (
	"context"
	"fmt"
)

type ShutdownWaiter interface {
	// Implementors of WaitUntilShutdown are expected to block until any goroutine they spawned have finished, or until the provided context has closed
	// Any persistent goroutine started with Forever is a ShutdownWaiter
	WaitUntilShutdown(shutdownContext context.Context)
}

// ChanShutdownWaiter need more description.

type ChanShutdownWaiter struct {
	closed      chan struct{}
	description string
}

func NewChanWaiter(description string) ChanShutdownWaiter {
	return ChanShutdownWaiter{closed: make(chan struct{}), description: description}
}

func (c *ChanShutdownWaiter) WaitUntilShutdown(shutdownContext context.Context) {
	select {
	case <-c.closed:
	case <-shutdownContext.Done():
		if shutdownContext.Err() == context.DeadlineExceeded {
			panic(fmt.Sprintf("failed to shutdown %s before timeout", c.description))
		}
	}
}

func (c *ChanShutdownWaiter) Shutdown() {
	close(c.closed)
}
