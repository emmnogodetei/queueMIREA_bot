package handlers

import (
	"context"

	"github.com/go-telegram/bot"
)

func IsAdmin(ctx context.Context, b *bot.Bot, chatID, userID int64) bool{
	admins, err := b.GetChatAdministrators(ctx, &bot.GetChatAdministratorsParams{
		ChatID: chatID,
	})

	if err != nil{
		return false
	}

	for _, admin := range admins{
		if (admin.Owner != nil && admin.Owner.User.ID == userID) || (admin.Administrator != nil && admin.Administrator.User.ID == userID) {
			return true
		}
	}
	return false
}