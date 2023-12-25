package factorysetting

import "github.com/suzuito/sandbox2-go/crawler/internal/usecase/discord"

type NotifierFactorySetting struct {
	DiscordClient discord.DiscordGoSession
}
