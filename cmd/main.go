package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/theerSs/vtt/internal/platform/app"
	"github.com/theerSs/vtt/internal/platform/env"
)

func main() {
	if err := env.Load(); err != nil {
		log.Fatalf("env load failed: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
