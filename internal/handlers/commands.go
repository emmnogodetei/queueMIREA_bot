package handlers

import (
	"context"
	"math/rand"
	"fmt"
	"slices"
	"strings"
	

	"github.com/emmnogodetei/queueMIREA_bot/internal/queue"
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

func CreateQueue(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	text := "The queue has already been created"
	if queue.ChatData[update.Message.Chat.ID] == nil {
		queue.ChatData[update.Message.Chat.ID] = make([]string, 0, queue.GroupSize)
		text = "Queue has been created. Now click + or +&#60;nubmer&#62; to join the queue"
	}
	_, err := tgbot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageThreadID: update.Message.MessageThreadID,
		Text:   text,
	})
	if err != nil {
		fmt.Printf("error sending message: %v\n", err)
		return
	}
}

func DeleteQueue(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	text := "Yod didn't create a queue"
	if queue.ChatData[update.Message.Chat.ID] != nil {
		delete(queue.ChatData, update.Message.Chat.ID)
		text = "Queue has been deleted"
	}
	_, err := tgbot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageThreadID: update.Message.MessageThreadID,
		Text:   text,
	})
	if err != nil {
		fmt.Printf("error sending message: %v\n", err)
		return
	}
}

func AddToQueue(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	if queue.ChatData[update.Message.Chat.ID] == nil {
		return
	}
	FLName := update.Message.From.FirstName + " " + update.Message.From.LastName
	if !slices.Contains(queue.ChatData[update.Message.Chat.ID], FLName) {
		queue.ChatData[update.Message.Chat.ID] = append(queue.ChatData[update.Message.Chat.ID], FLName)
	}
}

func GetQueue(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	text := "You didn't create a queue"
	if queue.ChatData[update.Message.Chat.ID] != nil && len(queue.ChatData[update.Message.Chat.ID]) != 0 {
		text = "Queue:\n\n" + strings.Join(queue.ChatData[update.Message.Chat.ID], "\n")
	}
	_, err := tgbot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageThreadID: update.Message.MessageThreadID,
		Text:   text,
	})
	if err != nil {
		fmt.Printf("error sending message: %v\n", err)
		return
	}
}

func ShuffleQueue(ctx context.Context, tgbot *bot.Bot, update *models.Update){
	text := "You didn't create a queue"
	if queue.ChatData[update.Message.Chat.ID] != nil && len(queue.ChatData[update.Message.Chat.ID]) != 0 {
		randomQueue := make([]string,len(queue.ChatData[update.Message.Chat.ID]))
		copy(randomQueue,queue.ChatData[update.Message.Chat.ID])
		for i := range randomQueue {
			j := rand.Intn(i + 1)
			randomQueue[i], randomQueue[j] = randomQueue[j], randomQueue[i]
		}
		text = "Random queue:\n\n" + strings.Join(randomQueue, "\n")
	}

	_, err := tgbot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageThreadID: update.Message.MessageThreadID,
		Text:   text,
	})
	if err != nil {
		fmt.Printf("error sending message: %v\n", err)
		return
	}
}
