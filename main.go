package main

import (
	"github.com/fevral13/twigo/handlers"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Starting the server")

	router := http.NewServeMux()

	router.Handle("GET /", handlers.NewIndex())

	err := http.ListenAndServe(":9001", router)
	if err != nil {
		logger.Error("aaaaaa", err)
	}
}
