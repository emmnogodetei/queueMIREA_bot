package bot

import (
	"context"

	"github.com/emmnogodetei/queueMIREA_bot/filters"
	"github.com/emmnogodetei/queueMIREA_bot/handlers"
	"github.com/go-telegram/bot"
)

func Run(ctx context.Context, token string) error {
	b, err := bot.New(token)
	if err!=nil{
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.Start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "+", bot.MatchTypeExact, handlers.AddDefault)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/remove", bot.MatchTypeExact, handlers.RemoveQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get", bot.MatchTypeExact, handlers.GetQueue)

	b.RegisterHandlerMatchFunc(filters.IsPlusWithPriority, handlers.AddWithPriority)

	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "btn_", bot.MatchTypePrefix, handlers.CallbackHandler)

	b.Start(ctx)
	return nil
}