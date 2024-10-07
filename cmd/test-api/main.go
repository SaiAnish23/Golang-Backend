package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SaiAnish23/Golang-Backend/internal/config"
	"github.com/SaiAnish23/Golang-Backend/internal/http/handlers/test"
)

func main() {

	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/test", test.New())

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	slog.Info("server is starting")
	fmt.Println("server is running on port", cfg.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}

	}()

	<-done

	slog.Info("server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("server shutdown failed")
	}

	slog.Info("server shutdown completed")

}
