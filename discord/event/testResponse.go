package event

import "github.com/bwmarrin/discordgo"

// FIXME: 작동 테스트를 위한 임시 함수
func SimpleResponse(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == "test" {
		discord.ChannelMessageSend(message.ChannelID, "It's working")
		return
	}
	return
}
