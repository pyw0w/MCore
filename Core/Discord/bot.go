package Discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID   string
	Session *discordgo.Session
)

func Connect(token string) {
	// Create a new Discord session using the provided bot token.
	var err error
	Session, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := Session.User("@me")
	if err != nil {
		log.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	log.Println("Bot connected")
}

func Start() {
	// Open the websocket and begin listening.
	err := Session.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	return
}

func AddHandler(handler interface{}) {
	Session.AddHandler(handler)
}
