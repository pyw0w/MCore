package Discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	_       string
	Session *discordgo.Session
)

func Connect(token string) {
	// Создаем новую сессию Discord с использованием предоставленного токена бота.
	var err error
	Session, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Println("Ошибка при создании сессии Discord:", err)
		return
	}

	// Получаем информацию об аккаунте.
	u, err := Session.User("@me")
	if err != nil {
		log.Println("Ошибка при получении данных об аккаунте:", err)
	}

	// Сохраняем ID аккаунта для дальнейшего использования.
	_ = u.ID
}

func Start() {
	// Открываем вебсокет и начинаем прослушивание.
	err := Session.Open()
	if err != nil {
		log.Println("Ошибка при открытии соединения:", err)
		return
	}

	log.Println("Бот работает. Нажмите CTRL-C для выхода.")

	return
}

func AddHandler(handler interface{}) {
	Session.AddHandler(handler)
}
