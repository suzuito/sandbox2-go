package govillagenote

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
	"go.uber.org/mock/gomock"
)

func TestNotifier(t *testing.T) {
	n := Notifier{}
	assert.Equal(t, notifier.NotifierID("govillagenote"), n.ID())
	assert.IsType(t, &note.TimeSeriesDataNoteArticle{}, n.NewEmptyTimeSeriesData(context.Background()))
}

func TestNotify(t *testing.T) {
	testCases := []struct {
		desc        string
		setUp       func(mockDiscord *discord.MockDiscordGoSession)
		inputD      timeseriesdata.TimeSeriesData
		expectedErr string
	}{
		{
			desc: "Success",
			inputD: &note.TimeSeriesDataNoteArticle{
				URL:            "https://www.example.com",
				Title:          "title1",
				Description:    "desc1",
				ArticleContent: "cont1",
				ImageURL:       "https://www.example.com/1",
				PublishedAt:    time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
				Tags: []note.TimeSeriesDataNoteArticleTag{
					{Name: "tag1"},
				},
			},
			setUp: func(mockDiscord *discord.MockDiscordGoSession) {
				mockDiscord.EXPECT().ChannelMessageSendComplex(
					"dummyChannelID",
					gomock.Any(),
				)
			},
		},
		{
			desc:   "Error if inputD does not type 'note.TimeSeriesDataNoteArticle'",
			inputD: &timeseriesdata.TimeSeriesDataBlogFeed{},
			setUp: func(mockDiscord *discord.MockDiscordGoSession) {
			},
			expectedErr: `Cannot cast from timeseriesdata\.TimeSeriesData to \*note\.TimeSeriesDataNoteArticle`,
		},
		{
			desc: "Error if ChannelMessageSendComplex return error",
			inputD: &note.TimeSeriesDataNoteArticle{
				URL:            "https://www.example.com",
				Title:          "title1",
				Description:    "desc1",
				ArticleContent: "cont1",
				ImageURL:       "https://www.example.com/1",
				PublishedAt:    time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
				Tags: []note.TimeSeriesDataNoteArticleTag{
					{Name: "tag1"},
				},
			},
			setUp: func(mockDiscord *discord.MockDiscordGoSession) {
				mockDiscord.EXPECT().ChannelMessageSendComplex(
					"dummyChannelID",
					gomock.Any(),
				).Return(nil, errors.New("dummy error"))
			},
			expectedErr: "Cannot send to discord channel : dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sess := discord.NewMockDiscordGoSession(ctrl)
			tC.setUp(sess)
			n := NewNotifier(sess, "dummyChannelID")
			err := n.Notify(ctx, tC.inputD)
			test_helper.AssertError(t, tC.expectedErr, err)
		})
	}
}
