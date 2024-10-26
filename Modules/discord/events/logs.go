package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func OnMessage(s *discordgo.Session, event *discordgo.Message) {
	log.Printf(event.Author.Username, event.Content)
}
