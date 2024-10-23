package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	DBot "MiniCore/Core/Discord"
	"github.com/joho/godotenv"
)

var (
	discordToken  string
	telegramToken string
)

var version = "undefined"

func init() {
	log.Println("Starting Server Status " + version)
}

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	// Получаем токены из окружения
	discordToken = os.Getenv("DISCORD_TOKEN")
	telegramToken = os.Getenv("TELEGRAM_TOKEN")

	if discordToken == "" {
		log.Fatalf("Токен Discord не найдены в .env файле")
	} else {
		DBot.Connect(discordToken)
	}

	if telegramToken == "" {
		log.Fatalf("Токен Telegram не найдены в .env файле")
	} else {

	}

	// Обработка CTRL + C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGINT, os.Interrupt)
	<-stop

	log.Println("Завершена")
}
