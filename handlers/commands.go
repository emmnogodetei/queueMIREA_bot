// Package handlers содержит обработчики Telegram команд
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

// Start присылает приветственное сообщение
func Start(ctx context.Context, tgbot *bot.Bot, update *models.Update) {
	tgbot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageThreadID: update.Message.MessageThreadID,
		Text:   "Привет! Я бот для формирования очередей в чатах",
	})
}
// Help выводит справку
func Help(ctx context.Context, tgbot *bot.Bot, update *models.Update){
	text := "Нажмите (+), чтобы присоединиться к очереди\n"
	text+= "Нажмите (+<число>), чтобы присоединиться к очереди с приоритетом <число>\n"
	text+= "Нажмите /remove (только для администраторов), чтобы удалить список очерди в чате или в топике\n"
	text+= "Нажмите /get для получения списке очереди с встроенной клавиатурой\n\n"
	text+= "Кнопка 'Обновить' обновляет список очереди\n"
	text+= "Кнопка 'Отсортировать' (только для администраторов) сортирует очередь по приоритету; чем ниже число приритета, тем выше человек в очереди\n"
	text+="Кнопка 'Удалить вершину' (только для администраторов) удаляет первую строку в очереди\n"
	text+="Кнопка 'Удалить себя' удаляет пользователя, нажавшего на кнопку из списка очереди\n"
	text+="Кнопка 'Удалить это сообщение' удалет сообщение"
}

// AddDefault обрабатывает команду добавления в очередь
//
// Триггерится сообщением "+"
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

// AddWithPriority обрабатывает команду добавления в очередь
//
// Триггерится сообщением "+<число>"
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

// RemoveQueue удаляет список очереди внутри одного чата и внутри одного топика, если есть
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

// GetQueue возвращает список очереди с inline клавиатурой
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

