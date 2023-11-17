package discordconnpassevent

import (
	"context"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/crawler/pkg/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Notifier struct {
	DiscordClient    discord.DiscordGoSession
	DiscordChannelID string
}

func (t *Notifier) ID() notifier.NotifierID {
	return "discordconnpassevent"
}

func (t *Notifier) NewEmptyTimeSeriesData() timeseriesdata.TimeSeriesData {
	return &timeseriesdata.TimeSeriesDataConnpassEvent{}
}

func (t *Notifier) Notify(
	ctx context.Context,
	d timeseriesdata.TimeSeriesData,
) error {
	event, ok := d.(*timeseriesdata.TimeSeriesDataConnpassEvent)
	if !ok {
		return terrors.Wrapf("Cannot cast from timeseriesdata.TimeSeriesData to *timeseriesdata.TimeSeriesDataConnpassEvent")
	}
	colorCode, _ := strconv.ParseInt("FF0000", 16, 64)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	if _, err := t.DiscordClient.ChannelMessageSendComplex(
		t.DiscordChannelID,
		&discordgo.MessageSend{
			Content: "",
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       event.Title,
					URL:         event.EventURL,
					Description: event.Catch,
					Color:       int(colorCode),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "connpass",
						IconURL: "https://connpass.com/static/img/api/connpass_logo_4.png",
					},
					Fields: []*discordgo.MessageEmbedField{
						{Name: "開始", Value: event.StartedAt.In(loc).Format(time.RFC3339), Inline: true},
						{Name: "終了", Value: event.EndedAt.In(loc).Format(time.RFC3339), Inline: true},
						{Name: "開催場所", Value: event.Place, Inline: true},
					},
				},
			},
		},
	); err != nil {
		return terrors.Wrapf("Cannot send to discord channel : %+v", err)
	}
	return nil
}

func New(def *notifier.NotifierDefinition, arg *factory.NewFuncNotifierArgument) (notifier.Notifier, error) {
	n := Notifier{
		DiscordClient: arg.DiscordClient,
	}
	if def.ID != n.ID() {
		return nil, factory.ErrNoMatchedNotifierID
	}
	discordChannelID, err := argument.GetFromArgumentDefinition[string](def.Argument, "DiscordChannelID")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	n.DiscordChannelID = discordChannelID
	return &n, nil
}
