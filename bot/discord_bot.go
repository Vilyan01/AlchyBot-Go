package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/Vilyan01/AlchyBot-Go/boterrors"
	"os"
	"fmt"
)

var (
	BotToken string
	BotID string
)

type DiscordBot struct {
	dg *discordgo.Session
	user *discordgo.User
}

func init() {
	BotToken = os.Getenv("DISCORDBOTTOKEN")
}

func NewBot() (DiscordBot,error) {
	var d DiscordBot
	// Check to see if the bot token has been created.
	if BotToken == "" {
		return d, boterrors.NILDISCORDTOKEN
	}

	// Create the bot using the token
	b, err := discordgo.New("Bot "+BotToken)
	if err != nil {
		return d, err
	}

	// Get the bot user
	u, err := b.User("@me")
	if err != nil {
		return d, err
	}
	BotID = u.ID

	// Create the discord bot struct with all the info.
	d = DiscordBot{dg: b, user: u}
	return d, err
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_,_ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if m.Content == "Ding" {
		_,_ = s.ChannelMessageSend(m.ChannelID, "Dang")
	}
}

func (d DiscordBot) Run() error {
	var e error
	d.dg.AddHandler(messageHandler)
	e = d.dg.Open()
	if e != nil {
		return e
	}
	fmt.Println("Bot is running...")
	<-make(chan struct{})
	return e
}
