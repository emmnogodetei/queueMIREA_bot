package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/emmnogodetei/queueMIREA_bot/storage"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func buildKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Удалить вершину", CallbackData: "btn_pop"},
				{Text: "Удалить себя", CallbackData: "btn_removeMe"},
				{Text: "Удалить это сообщение", CallbackData: "btn_delete"},
			},
		},
	}

	return kb
}

func CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// answering callback query first to let Telegram know that we received the callback query,
	// and we're handling it. Otherwise, Telegram might retry sending the update repetitively
	// as it thinks the callback query doesn't reach to our application. learn more by
	// reading the footnote of the https://core.telegram.org/bots/api#callbackquery type.
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	var err error
	var queue []string
	switch update.CallbackQuery.Data {
	case "btn_pop":
		err = storage.Pop(
			update.CallbackQuery.Message.Message.Chat.ID,
			int64(update.CallbackQuery.Message.Message.MessageThreadID),
		)
	case "btn_removeMe":
		err = storage.RemovePersone(
			update.CallbackQuery.Message.Message.Chat.ID,
			int64(update.CallbackQuery.Message.Message.MessageThreadID),
			update.CallbackQuery.From.ID,
		)
	case "btn_delete":
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.Message.ID,
		})
		return
	}
	if err != nil{
		fmt.Printf("error: %v\n", err)
	}

	queue, err = storage.Get(
		update.CallbackQuery.Message.Message.Chat.ID,
		int64(update.CallbackQuery.Message.Message.MessageThreadID),
	)

	if err != nil{
		fmt.Printf("error getting queue: %v\n", err)
		return
	}

	text := "the queue is empty"
	if len(queue)!=0{
		text = strings.Join(queue,"\n")
	}

	_, err = b.EditMessageText(ctx,&bot.EditMessageTextParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
		Text: text,
		ReplyMarkup: buildKeyboard(),
	})

	if err != nil{
		fmt.Printf("error sending message: %v\n", err)
	}
}

