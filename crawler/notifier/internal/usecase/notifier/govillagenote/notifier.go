package govillagenote

import (
	"context"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

const NotifierID notifier.NotifierID = "govillagenote"

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
	return &note.TimeSeriesDataNoteArticle{}
}

func (t *Notifier) Notify(
	ctx context.Context,
	d timeseriesdata.TimeSeriesData,
) error {
	noteFeed, ok := d.(*note.TimeSeriesDataNoteArticle)
	if !ok {
		return terrors.Wrapf("Cannot cast from timeseriesdata.TimeSeriesData to *note.TimeSeriesDataNoteArticle")
	}
	colorCode, _ := strconv.ParseInt("00FF00", 16, 64)
	if _, err := t.dcli.ChannelMessageSendComplex(
		t.discordChannelID,
		&discordgo.MessageSend{
			Content: "",
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       noteFeed.Title,
					URL:         noteFeed.URL,
					Description: noteFeed.Description,
					Image: &discordgo.MessageEmbedImage{
						URL: noteFeed.ImageURL,
					},
					Color: int(colorCode),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "note.com",
						IconURL: "https://d2l930y2yx77uc.cloudfront.net/assets/e_note_logo_202212-dd631ee36da01c6df532a635c43a845529ab68a43ff7b827ce39e965eb298b48.png",
					},
					// TODO 関連Orgをここに
					// Author: &discordgo.MessageEmbedAuthor{
					// 	Name: "hoge",
					// 	URL:  "https://www.example.com/",
					// },
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
