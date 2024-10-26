package commands

import (
	"github.com/bwmarrin/discordgo"
)

type ICommand interface {
	Execute(session *discordgo.Session, interaction *discordgo.InteractionCreate)
}

type Command struct {
	data discordgo.ApplicationCommand
}

var All = []ICommand{
	&Hello,
}
