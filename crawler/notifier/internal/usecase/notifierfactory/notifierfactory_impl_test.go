package notifierfactory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/usecase/notifier/govillageblogfeed"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/usecase/notifier/govillageconnpassevent"
	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/usecase/notifier/govillagegolangweekly"
)

func TestNotifierFactoryImpl(t *testing.T) {
	factory := NewDefaultNotifierFactoryImpl(
		nil,
		"",
		"",
	)
	testCases := []struct {
		desc                string
		inputFullPath       string
		expectedNotifierIDs []notifier.NotifierID
		expectedError       string
	}{
		{
			inputFullPath: "Crawler/TimeSeriesData/goblog/doc1",
			expectedNotifierIDs: []notifier.NotifierID{
				govillageblogfeed.NotifierID,
			},
		},
		{
			inputFullPath: "Crawler/TimeSeriesData/goconnpass/doc1",
			expectedNotifierIDs: []notifier.NotifierID{
				govillageconnpassevent.NotifierID,
			},
		},
		{
			inputFullPath: "Crawler/TimeSeriesData/golangweekly/doc1",
			expectedNotifierIDs: []notifier.NotifierID{
				govillagegolangweekly.NotifierID,
			},
		},
	}
	for _, tC := range testCases {
		desc := tC.inputFullPath
		if tC.desc != "" {
			desc = tC.desc
		}
		t.Run(desc, func(t *testing.T) {
			notifiers, err := factory.GetNotiferFromDocPathFirestore(context.Background(), tC.inputFullPath)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				notifierIDs := []notifier.NotifierID{}
				for _, n := range notifiers {
					notifierIDs = append(notifierIDs, n.ID())
				}
				for _, expectedNotifierID := range tC.expectedNotifierIDs {
					assert.Contains(t, notifierIDs, expectedNotifierID)
				}
			}
		})
	}
}
