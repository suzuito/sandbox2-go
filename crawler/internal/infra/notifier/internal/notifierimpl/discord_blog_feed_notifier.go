package notifierimpl

import (
	"context"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type DiscordBlogFeedNotifier struct {
	DiscordClient    discord.DiscordGoSession
	DiscordChannelID string
}

func (t *DiscordBlogFeedNotifier) ID() notifier.NotifierID {
	return "discordblogfeed"
}

func (t *DiscordBlogFeedNotifier) NewEmptyTimeSeriesData() timeseriesdata.TimeSeriesData {
	return &timeseriesdata.TimeSeriesDataBlogFeed{}
}

func (t *DiscordBlogFeedNotifier) Notify(
	ctx context.Context,
	d timeseriesdata.TimeSeriesData,
) error {
	blogFeed, ok := d.(*timeseriesdata.TimeSeriesDataBlogFeed)
	if !ok {
		return terrors.Wrapf("Cannot cast from timeseriesdata.TimeSeriesData to *timeseriesdata.TimeSeriesDataBlogFeed")
	}
	colorCode, _ := strconv.ParseInt("0000FF", 16, 64)
	// loc, _ := time.LoadLocation("Asia/Tokyo")
	embed := discordgo.MessageEmbed{}
	embed.Title = blogFeed.Title
	embed.Description = blogFeed.Summary
	embed.URL = blogFeed.URL
	embed.Color = int(colorCode)
	if blogFeed.Thumbnail != nil {
		embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL: blogFeed.Thumbnail.ImageURL,
		}
	}
	if blogFeed.Author != nil {
		embed.Author = &discordgo.MessageEmbedAuthor{
			Name:    blogFeed.Author.Name,
			URL:     blogFeed.Author.URL,
			IconURL: blogFeed.Author.ImageURL,
		}
	}
	if _, err := t.DiscordClient.ChannelMessageSendComplex(
		t.DiscordChannelID,
		&discordgo.MessageSend{
			Content: "",
			Embeds: []*discordgo.MessageEmbed{
				&embed,
			},
		},
	); err != nil {
		return terrors.Wrapf("Cannot send to discord channel : %+v", err)
	}
	return nil
}
