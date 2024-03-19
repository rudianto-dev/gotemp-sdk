package gracefully

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func WaitAndShutdown(ops map[string]any) {
	idleConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)
		ss := <-sigint
		log.Printf("received signal %s\n", ss.String())

		var wg sync.WaitGroup
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				if val, ok := innerOp.(func() error); ok {
					err := val()
					if err != nil {
						log.Printf("%s error while closing: %s", innerKey, err.Error())
					}
				} else if val, ok := innerOp.(func(ctx context.Context) error); ok {
					err := val(context.Background())
					if err != nil {
						log.Printf("%serror while closing: %s", innerKey, err.Error())
					}
				} else if val, ok := innerOp.(func()); ok {
					val()
				} else if val, ok := innerOp.(func(ctx context.Context)); ok {
					val(context.Background())
				}

				log.Printf("%s was shutdown", innerKey)
			}()
		}
		wg.Wait()
		close(idleConnClosed)
	}()

	<-idleConnClosed
	log.Printf("all system has been shutdown")
}
