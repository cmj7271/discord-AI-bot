package main

import (
	"log"

	"github.com/cmj7271/discord-AI-bot.git/discord"
)

func main() {
	if err := run(); err != nil {
		log.Printf("fail to start: %v\n", err)
	}
}

// TODO: context 추가해서 error 추가 구현
// TODO: PORT 번호는 환경변수로?
func run() error {
	err := discord.Run()
	if err != nil {
		log.Fatal("error runing bot: \n" + err.Error())
	}
	return nil
}
