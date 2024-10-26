package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func OnReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("Бот успешно подключен и готов к работе!")
	log.Printf("Зашел под именем %s#%s\n", event.User.Username, event.User.Discriminator)
}
