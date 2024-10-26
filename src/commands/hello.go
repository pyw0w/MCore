package commands

import "github.com/bwmarrin/discordgo"

type HelloCommand struct {
	Command
}

var Hello = HelloCommand{
	Command{
		data: discordgo.ApplicationCommand{
			Name:        "hello",
			Description: "Hello world!",
		},
	},
}

func (hello *HelloCommand) Execute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
}
