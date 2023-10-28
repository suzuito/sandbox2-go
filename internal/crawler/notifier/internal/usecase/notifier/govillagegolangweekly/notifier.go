package govillagegolangweekly

import (
	"context"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

const NotifierID notifier.NotifierID = "govillagegolangweekly"

type Notifier struct {
	dcli             discord.DiscordGoSession
	discordChannelID string
}

func (t *Notifier) ID() notifier.NotifierID {
	return NotifierID
}

func (t *Notifier) NewEmptyTimeSeriesData(
	ctx context.Context,
) timeseriesdata.TimeSeriesData {
	return &timeseriesdata.TimeSeriesDataBlogFeed{}
}

func (t *Notifier) Notify(
	ctx context.Context,
	d timeseriesdata.TimeSeriesData,
) error {
	blogFeed, ok := d.(*timeseriesdata.TimeSeriesDataBlogFeed)
	if !ok {
		return terrors.Wrapf("Cannot cast from timeseriesdata.TimeSeriesData to *timeseriesdata.TimeSeriesDataBlogFeed")
	}
	clog.L.Infof(ctx, "Notify %+v", blogFeed)
	colorCode, _ := strconv.ParseInt("00FF00", 16, 64)
	if _, err := t.dcli.ChannelMessageSendComplex(
		t.discordChannelID,
		&discordgo.MessageSend{
			Content: "Golang Weekly",
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:     blogFeed.Title,
					URL:       blogFeed.URL,
					Timestamp: blogFeed.PublishedAt.Format(time.RFC3339),
					Color:     int(colorCode),
				},
			},
		},
	); err != nil {
		return terrors.Wrapf("Cannot send to discord channel : %+v", err)
	}
	return nil
}

func NewNotifier(
	dcli discord.DiscordGoSession,
	discordChannelID string,
) *Notifier {
	return &Notifier{
		dcli:             dcli,
		discordChannelID: discordChannelID,
	}
}
