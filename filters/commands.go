// Package filters содержит фильтры сообщений

package filters

import (
	"fmt"
	"strconv"

	"github.com/go-telegram/bot/models"
)

// IsPlusWithPriority проверяет соответствует ли введенное сообщение формату (+)<число>
func IsPlusWithPriority(update *models.Update) bool{
	if update.Message == nil{
		return false
	}
	msg := update.Message.Text
	if len(msg) <=1{
		return false
	}

	num ,err := strconv.Atoi(msg[1:])
	if err != nil{
		fmt.Printf("error getting queue: %v\n", err)
		return false
	}
	return string(msg[0]) == "+" && num >= 1
}