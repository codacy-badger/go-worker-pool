package main

import (
	"errors"
	"github.com/dalmarcogd/go-worker-pool/server"
	"github.com/dalmarcogd/go-worker-pool/worker"
	"log"
	"time"
)

func main() {
	if err := server.
		New().
		Stats().
		HealthCheck().
		HandleError(func(w *worker.Worker, err error) {
			log.Printf("Worker [%s] error: %s", w.Name, err)
		}).
		Worker(
			"w1",
			func() error {
				time.Sleep(10 * time.Second)
				return errors.New("teste")
			},
			5).
		Worker(
			"w2",
			func() error {
				time.Sleep(30 * time.Second)
				return nil
			},
			3).
		Worker(
			"w3",
			func() error {
				time.Sleep(1 * time.Minute)
				return errors.New("teste")
			},
			2).
		//Worker(
		//	"w4",
		//	func() error {
		//		time.Sleep(1000)
		//		return nil
		//	},
		//	1).
		Run(); err != nil {
		panic(err)
	}
}
