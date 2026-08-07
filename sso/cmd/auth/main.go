package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/fa4ar/authsys/internal/config"
)

const envLocal = "LOCAL"

func main() {
	cfg := config.MustLoadConfig()

	fmt.Printf("%#v\n", cfg)

}

func SetUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
}
