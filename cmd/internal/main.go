package main

import (
	"context"
	"log"
	"mail_sender/cmd/internal/app"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}

}
