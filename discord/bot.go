package discord

import (
	"errors"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v11"
)

type config struct {
	token string
}

func Run() error {
	var cfg config
	// env 불러오기
	if err := env.Parse(&cfg); err != nil {
		return errors.New("env loading fatal: " + err.Error())
	}

	// discord 봇 생성
	discord, err := discordgo.New("Bot " + cfg.token)
	if err != nil {
		return errors.New("error in connecting: " + err.Error())
	}

	// TODO: add a event handler
	// discord.AddHandler()

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
