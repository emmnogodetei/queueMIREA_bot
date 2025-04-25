package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/emmnogodetei/queueMIREA_bot/storage"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)


func Start(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	tgbot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageThreadID: update.Message.MessageThreadID,
		Text:   "Hi! I am a bot for forming queues in chats.",
	})
}

func AddDefault(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	FLName := update.Message.From.FirstName + " " + update.Message.From.LastName 
	
	err := storage.Add(
		update.Message.Chat.ID,
		int64(update.Message.MessageThreadID),
		update.Message.From.ID,
		FLName,
		update.Message.From.Username,
		0,
	)
	
	if err != nil {
		fmt.Printf("error adding to queue: %v\n", err)
	}
}

func AddWithPriority(ctx context.Context, tgbot *bot.Bot, update *models.Update){
	FLName := update.Message.From.FirstName + " " + update.Message.From.LastName 
	num, _ := strconv.Atoi(update.Message.Text[1:])

	err := storage.Add(
		update.Message.Chat.ID,
		int64(update.Message.MessageThreadID),
		update.Message.From.ID,
		FLName,
		update.Message.From.Username,
		num,
	)
	
	if err != nil {
		fmt.Printf("error adding to queue: %v\n", err)
	}
}

func RemoveQueue(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	if !IsAdmin(ctx,tgbot, update.Message.Chat.ID, update.Message.From.ID){
		return
	}
	err := storage.Remove(
		update.Message.Chat.ID,
		int64(update.Message.MessageThreadID),
	)

	if err != nil{
		fmt.Printf("error deleting the queue: %v\n", err)
	}
}

func GetQueue(ctx context.Context, tgbot *bot.Bot, update *models.Update){
	queue, err := storage.Get(
		update.Message.Chat.ID,
		int64(update.Message.MessageThreadID),
	)

	if err != nil{
		fmt.Printf("error getting queue: %v\n", err)
		return
	}
	text := "the queue is empty"
	if len(queue)!=0{
		text = strings.Join(queue,"\n")
	}
	_, err = tgbot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageThreadID: update.Message.MessageThreadID,
		Text: text,
		ReplyMarkup: buildKeyboard(update.Message.Chat.ID, int64(update.Message.MessageThreadID)),
	})
	if err != nil{
		fmt.Printf("error sending message: %v\n", err)
	}
}

