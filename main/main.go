package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	discordToken  string
	telegramToken string
)

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
		DSBot()
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

func DSBot() {
	dg, err := discordgo.New("Bot " + discordToken)

	// Информация статуса бота
	dg.AddHandler(func(dg *discordgo.Session, r *discordgo.Ready) {
		log.Println("Бот запустился")
	})
	// Обработка ошибок сессии
	if err != nil {
		log.Fatal("Ошибка при создании сессии Discord: ", err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Fatal("Ошибка при подключении к Discord: ", err)
	}
}
