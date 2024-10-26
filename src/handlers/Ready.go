package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func Ready() {
	Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
}
