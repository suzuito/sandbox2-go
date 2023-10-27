package govillageconnpassevent

import (
	"context"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

const NotifierID notifier.NotifierID = "govillageconnpassevent"

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
	clog.L.Infof(ctx, "Notify %+v", event)
	colorCode, _ := strconv.ParseInt("FF0000", 16, 64)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	if _, err := t.dcli.ChannelMessageSendComplex(
		t.discordChannelID,
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

func NewNotifier(
	dcli discord.DiscordGoSession,
	discordChannelID string,
) *Notifier {
	return &Notifier{
		dcli:             dcli,
		discordChannelID: discordChannelID,
	}
}
