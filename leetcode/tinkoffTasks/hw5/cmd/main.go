package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tinkoffTasks/hw5/internal/handlers"
	"tinkoffTasks/hw5/internal/repositories"
	"tinkoffTasks/hw5/internal/services"
)

func newServer(addr string, logger *logrus.Logger) *http.Server {
	usersRepo := repositories.NewUsersRepository(logger)
	usersService := services.NewUsersService(usersRepo, logger)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	userHandler := handlers.NewUsersHandler(usersService, logger)
	r.Mount("/users", userHandler.Routes())

	return &http.Server{
		Handler: r,
		Addr:    addr,
	}
}

func execute(addr string, logger *logrus.Logger) (<-chan error, error) {
	srv := newServer(addr, logger)

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
	logger := logrus.New()

	errC, err := execute(":5000", logger)
	if err != nil {
		log.Fatalln("Can't run")
	}
	if err := <-errC; err != nil {
		log.Fatalln("Error while execution")
	}
}
