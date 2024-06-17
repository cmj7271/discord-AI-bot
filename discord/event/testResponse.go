package event

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// FIXME: 작동 테스트를 위한 임시 함수
func SimpleResponse(discord *discordgo.Session, message *discordgo.MessageCreate) error {
	if message.Content == "test" {
		_, err := discord.ChannelMessageSend(message.ChannelID, "It's working")
		if err != nil {
			return errors.New("error: sending message:\n\t" + string(err.Error()))
		}
		return nil
	}
	return nil
}
