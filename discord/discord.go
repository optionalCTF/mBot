package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/un4gi/mBot/config"
)

var BotID string

func ConnectDiscord(m string) {
	d, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error starting new bot session,", err)
	}

	user, err := d.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = user.ID

	err = d.Open()
	if err != nil {
		fmt.Println("Error connecting to Discord.")
	}

	fmt.Println(fmt.Sprint(user), "has connected to Discord!")
	SendMessage(d, m)
}

func SendMessage(d *discordgo.Session, m string) {
	_, err := d.ChannelMessageSend(config.Channel, m)
	if err != nil {
		fmt.Println(err)
	}
}
