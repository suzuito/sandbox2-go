package inject

type Environment struct {
	GoVillageDiscordChannelIDNews   string `envconfig:"GO_VILLAGE_DISCORD_CHANNEL_ID_NEWS"`
	GoVillageDiscordChannelIDEvents string `envconfig:"GO_VILLAGE_DISCORD_CHANNEL_ID_EVENTS"`
	GoVillageDiscordBotToken        string `envconfig:"GO_VILLAGE_DISCORD_BOT_TOKEN"`
}
