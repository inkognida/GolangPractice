package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello"))
}

func newServer(addr string) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/hello", GetHello)

	return &http.Server{
		Handler: r,
		Addr:    addr,
	}
}

func execute(addr string) (<-chan error, error) {
	srv := newServer(addr)

	// Grateful shutdown
	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		log.Println("Grateful shutdown started")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		log.Println("Shutdown stopped")
	}()

	go func() {
		log.Println("Listen and serve")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, nil
}

func main() {
	errC, err := execute(":5000")
	if err != nil {
		log.Fatalln("Can't run")
	}
	if err := <-errC; err != nil {
		log.Fatalln("Error while execution")
	}
}
