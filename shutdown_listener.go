package govnr

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type OSShutdownListener struct {
	// Logger       log.Logger
	shutdownCond *sync.Cond
	shutdowner   GracefulShutdowner
}

// logger log.Logger,
func NewShutdownListener(shutdowner GracefulShutdowner) *OSShutdownListener {
	return &OSShutdownListener{
		shutdownCond: sync.NewCond(&sync.Mutex{}),
		// Logger:       logger,
		shutdowner: shutdowner,
	}
}

func (n *OSShutdownListener) ListenToOSShutdownSignal() {
	// if waiting for shutdown, listen for sigint and sigterm
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	Once(nil, func() {
		<-signalChan
		// n.Logger.Info("terminating node gracefully due to os signal received")
		ShutdownGracefully(n.shutdowner, 100*time.Millisecond)
	})

}
