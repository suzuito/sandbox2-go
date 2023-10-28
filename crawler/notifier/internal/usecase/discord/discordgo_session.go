package discord

import "github.com/bwmarrin/discordgo"

type DiscordGoSession interface {
	ChannelMessageSend(
		channelID string,
		data string,
		options ...discordgo.RequestOption,
	) (st *discordgo.Message, err error)
	ChannelMessageSendComplex(
		channelID string,
		data *discordgo.MessageSend,
		options ...discordgo.RequestOption,
	) (st *discordgo.Message, err error)
}
