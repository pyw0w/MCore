package telegram

import (
	TBot "MiniCore/Core/Telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Start(update tgbotapi.Update) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Отправляем простой ответ на команду /start
		if update.Message.Text == "/start" {
			TBot.SendMessage(update.Message.Chat.ID, "Привет! Это Telegram бот!")
		}
	}
}
