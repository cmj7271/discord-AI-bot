package discord

import (
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/cmj7271/discord-AI-bot.git/discord/event"
)

func Run(options ...SetOption) error {
	// env 불러오기
	var opts Options
	for _, setOpt := range options {
		err := setOpt(&opts)
		if err != nil {
			return err
		}
	}

	// discord 봇 생성
	fmt.Printf("%#v\n", opts.Token)
	discord, err := discordgo.New("Bot " + *opts.Token)
	if err != nil {
		return errors.New("error in connecting: " + err.Error())
	}

	// TODO: add a event handler
	discord.AddHandler(event.SimpleResponse)

	// Open creates a websocket connection to Discord.
	// https://pkg.go.dev/github.com/bwmarrin/discordgo#Session.Open
	err = discord.Open()
	if err != nil {
		return errors.New("error in open: " + err.Error())
	}
	defer discord.Close()

	// keep bot running
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	return nil
}
