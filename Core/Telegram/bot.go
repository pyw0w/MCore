package Telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Bot   *tgbotapi.BotAPI
	BotID int64
)

func Connect(token string) {
	var err error
	Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Ошибка при создании сессии Telegram:", err)
		return
	}

	// Получаем ID бота
	BotID = Bot.Self.ID

	log.Printf("Бот Telegram (%s) успешно подключен", Bot.Self.UserName)
}

func Start(handler func(update tgbotapi.Update)) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		handler(update)
	}
}

func SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка отправки сообщения:", err)
	}
}
