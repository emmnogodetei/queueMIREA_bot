package bot

import (
	"context"

	"github.com/emmnogodetei/queueMIREA_bot/internal/handlers"
	"github.com/go-telegram/bot"
)

func Run(ctx context.Context, token string) error {
	b, err := bot.New(token)
	if err!=nil{
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.Start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/create", bot.MatchTypeExact, handlers.CreateQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/delete", bot.MatchTypeExact, handlers.DeleteQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "+", bot.MatchTypeExact, handlers.AddToQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/queue", bot.MatchTypeExact, handlers.GetQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/random_queue", bot.MatchTypeExact, handlers.ShuffleQueue)

	b.Start(ctx)
	return nil
}