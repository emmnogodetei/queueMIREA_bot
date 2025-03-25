package main

import (
	"log"
	"os"
	"os/signal"
	"context"
	"github.com/emmnogodetei/queueMIREA_bot/internal/bot"
)

func main() {
	token := os.Getenv("TELEGRAM_API_TOKEN2")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := bot.Run(ctx, token); err != nil {
		log.Fatal(err)
	}
}