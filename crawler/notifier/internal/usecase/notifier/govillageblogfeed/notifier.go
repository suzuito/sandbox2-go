package govillageblogfeed

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

const NotifierID notifier.NotifierID = "govillageblogfeed"

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
	if _, err := t.dcli.ChannelMessageSend(t.discordChannelID, blogFeed.URL); err != nil {
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
