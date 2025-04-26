// Package bot содержит функцию функцию по запуску бота
package bot

import (
	"context"

	"github.com/emmnogodetei/queueMIREA_bot/filters"
	"github.com/emmnogodetei/queueMIREA_bot/handlers"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Run запускает бота, устанавливает список команд бота и регистрирует обработчиков
// Параметры:
//   - ctx: контекст выполнения
//   - token: API токен бота
func Run(ctx context.Context, token string) error {
	b, err := bot.New(token)
	if err!=nil{
		panic(err)
	}

	SetCommands(ctx, b)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.Start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "+", bot.MatchTypeExact, handlers.AddDefault)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/remove", bot.MatchTypeExact, handlers.RemoveQueue)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get", bot.MatchTypeExact, handlers.GetQueue)

	b.RegisterHandlerMatchFunc(filters.IsPlusWithPriority, handlers.AddWithPriority)

	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "btn_", bot.MatchTypePrefix, handlers.CallbackHandler)

	b.Start(ctx)
	return nil
}

// SetCommands  устанавливает список команд бота в чате
func SetCommands(ctx context.Context, b *bot.Bot) {
	commands := []models.BotCommand{
		{Command: "start", Description: "Запустить бота"},
		{Command: "help", Description: "Показать справку"},
		{Command: "get", Description: "Получить список очереди"},
		{Command: "remove", Description: "Удалить список очереди"},
	}

	_, err := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: commands,
		Scope:    &models.BotCommandScopeDefault{},
	})
	if err != nil {
		panic(err)
	}
}