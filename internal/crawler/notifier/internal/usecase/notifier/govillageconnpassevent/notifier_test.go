package govillageconnpassevent

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
	"go.uber.org/mock/gomock"
)

func TestNotifier(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sess := discord.NewMockDiscordGoSession(ctrl)
	n := NewNotifier(sess, "chan1")
	assert.Equal(t, notifier.NotifierID("govillageconnpassevent"), n.ID())
	assert.IsType(t, &timeseriesdata.TimeSeriesDataConnpassEvent{}, n.NewEmptyTimeSeriesData(ctx))
}

func TestNotify(t *testing.T) {
	testCases := []struct {
		desc                  string
		inputDiscordChannelID string
		inputD                timeseriesdata.TimeSeriesData
		expectedError         string
		setUp                 func(m *discord.MockDiscordGoSession)
	}{
		{
			desc:                  "success",
			inputDiscordChannelID: "chan1",
			inputD:                &timeseriesdata.TimeSeriesDataConnpassEvent{},
			setUp: func(m *discord.MockDiscordGoSession) {
				m.EXPECT().ChannelMessageSendComplex("chan1", gomock.Any())
			},
		},
		{
			desc:                  "invalid timeseriesdata",
			inputDiscordChannelID: "chan1",
			inputD:                &timeseriesdata.TimeSeriesDataBlogFeed{},
			expectedError:         `Cannot cast from timeseriesdata.TimeSeriesData to \*timeseriesdata.TimeSeriesDataConnpassEvent`,
			setUp:                 func(m *discord.MockDiscordGoSession) {},
		},
		{
			desc:                  "failed to ChannelMessageSend",
			inputDiscordChannelID: "chan1",
			inputD:                &timeseriesdata.TimeSeriesDataConnpassEvent{},
			expectedError:         "Cannot send to discord channel : dummy error",
			setUp: func(m *discord.MockDiscordGoSession) {
				m.EXPECT().
					ChannelMessageSendComplex("chan1", gomock.Any()).
					Return(nil, errors.New("dummy error"))
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := discord.NewMockDiscordGoSession(ctrl)
			tC.setUp(m)
			n := NewNotifier(m, tC.inputDiscordChannelID)
			err := n.Notify(context.Background(), tC.inputD)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
