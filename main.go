package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/emmnogodetei/queueMIREA_bot/bot"
	"github.com/emmnogodetei/queueMIREA_bot/storage"
)

func main() {
	token := os.Getenv("TELEGRAM_API_TOKEN2")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	storage.Init()

	if err := bot.Run(ctx, token); err != nil {
		log.Fatal(err)
	}
}