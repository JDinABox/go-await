package await

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Interrupt struct {
	*sync.WaitGroup
	closer chan struct{}
}

func NewInterrupt() *Interrupt {
	in := &Interrupt{
		WaitGroup: new(sync.WaitGroup),
		closer:    make(chan struct{}),
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		close(in.closer)
	}()

	return in
}

func (i *Interrupt) Await() {
	<-i.closer
}
