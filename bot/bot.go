package bot

import (
	"context"

	"github.com/emmnogodetei/queueMIREA_bot/handlers"
	"github.com/go-telegram/bot"
)

func Run(ctx context.Context, token string) error {
	b, err := bot.New(token)
	if err!=nil{
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.Start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "+", bot.MatchTypeExact, handlers.AddToQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/pop", bot.MatchTypeExact, handlers.RemoveFirst)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/remove_me", bot.MatchTypeExact, handlers.RemoveMe)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/remove", bot.MatchTypeExact, handlers.RemoveQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get", bot.MatchTypeExact, handlers.GetQueue)

	b.Start(ctx)
	return nil
}