package notifierimpl

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"go.uber.org/mock/gomock"
)

func TestDiscordConnpassEventNotifierNotifierID(t *testing.T) {
	n := &DiscordConnpassEventNotifier{}
	assert.Equal(t, notifier.NotifierID("discordconnpassevent"), n.ID())
}

func TestDiscordConnpassEventNotifierNewEmptyTimeSeriesData(t *testing.T) {
	n := &DiscordConnpassEventNotifier{}
	assert.IsType(t, &timeseriesdata.TimeSeriesDataEvent{}, n.NewEmptyTimeSeriesData())
}

func TestDiscordConnpassEventNotifierNotify(t *testing.T) {
	inputDiscordChannelID := "dummy"
	testCases := []struct {
		desc  string
		setUp func(
			mockDiscordClient *discord.MockDiscordGoSession,
		)
		inputD        timeseriesdata.TimeSeriesData
		expectedError string
	}{
		{
			desc: "Success",
			inputD: &timeseriesdata.TimeSeriesDataEvent{
				Organizer: &timeseriesdata.TimeSeriesDataEventOrganizer{},
			},
			setUp: func(
				mockDiscordClient *discord.MockDiscordGoSession,
			) {
				mockDiscordClient.EXPECT().ChannelMessageSendComplex(
					inputDiscordChannelID,
					gomock.Any(),
				).Return(nil, nil)
			},
		},
		{
			desc:   "Failed (not *timeseriesdata.TimeSeriesDataEvent)",
			inputD: &timeseriesdata.TimeSeriesDataBlogFeed{},
			setUp: func(
				mockDiscordClient *discord.MockDiscordGoSession,
			) {
			},
			expectedError: `Cannot cast from timeseriesdata.TimeSeriesData to \*timeseriesdata.TimeSeriesDataEvent`,
		},
		{
			desc:   "Failed (discord)",
			inputD: &timeseriesdata.TimeSeriesDataEvent{},
			setUp: func(
				mockDiscordClient *discord.MockDiscordGoSession,
			) {
				mockDiscordClient.EXPECT().ChannelMessageSendComplex(
					inputDiscordChannelID,
					gomock.Any(),
				).Return(nil, errors.New("dummy"))
			},
			expectedError: `dummy`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockDiscordClient := discord.NewMockDiscordGoSession(ctrl)
			tC.setUp(mockDiscordClient)
			notifier := &DiscordConnpassEventNotifier{
				DiscordChannelID: inputDiscordChannelID,
				DiscordClient:    mockDiscordClient,
			}
			err := notifier.Notify(context.Background(), tC.inputD)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
