package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"go.uber.org/mock/gomock"
)

func TestStartPipelinePeriodically(t *testing.T) {
	testCases := []struct {
		utTestCase
		inputCrawlerStarterSettingID crawler.CrawlerStarterSettingID
		expectedError                string
	}{
		{
			utTestCase: utTestCase{
				desc: "Success",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerStarterSettings(
							gomock.Any(),
							crawler.CrawlerStarterSettingID("starter001"),
						).
						Return([]*crawler.CrawlerStarterSetting{
							{
								ID:               "starter001",
								CrawlerID:        "crawler001",
								CrawlerInputData: crawler.CrawlerInputData{},
							},
						}, nil)
					mocks.MockTriggerCrawlerQueue.EXPECT().
						PublishCrawlEvent(
							context.Background(),
							crawler.CrawlerID("crawler001"),
							crawler.CrawlerInputData{},
						)
				},
				expectedLogLines: []string{
					"level=INFO msg=StartPipelinePeriodically crawlerStarterID=starter001",
					"level=INFO msg=PublishCrawlEvent crawlerStarterID=starter001 crawlerID=crawler001",
				},
			},
			inputCrawlerStarterSettingID: "starter001",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to GetCrawlerStarterSettings",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerStarterSettings(
							gomock.Any(),
							crawler.CrawlerStarterSettingID("starter001"),
						).
						Return(nil, errors.New("dummy"))
				},
				expectedLogLines: []string{
					"level=INFO msg=StartPipelinePeriodically crawlerStarterID=starter001",
					"level=ERROR msg=\"Failed to GetCrawlerStarterSettings\" crawlerStarterID=starter001 err=dummy",
				},
			},
			inputCrawlerStarterSettingID: "starter001",
			expectedError:                "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Skip failed to PublishCrawlEvent",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerStarterSettings(
							gomock.Any(),
							crawler.CrawlerStarterSettingID("starter001"),
						).
						Return([]*crawler.CrawlerStarterSetting{
							{
								ID:               "starter001",
								CrawlerID:        "crawler001",
								CrawlerInputData: crawler.CrawlerInputData{},
							},
							{
								ID:               "starter001",
								CrawlerID:        "crawler002",
								CrawlerInputData: crawler.CrawlerInputData{},
							},
						}, nil)
					gomock.InOrder(
						mocks.MockTriggerCrawlerQueue.EXPECT().
							PublishCrawlEvent(
								context.Background(),
								crawler.CrawlerID("crawler001"),
								crawler.CrawlerInputData{},
							).Return(errors.New("dummy")),
						mocks.MockTriggerCrawlerQueue.EXPECT().
							PublishCrawlEvent(
								context.Background(),
								crawler.CrawlerID("crawler002"),
								crawler.CrawlerInputData{},
							),
					)
				},
				expectedLogLines: []string{
					"level=INFO msg=StartPipelinePeriodically crawlerStarterID=starter001",
					"level=INFO msg=PublishCrawlEvent crawlerStarterID=starter001 crawlerID=crawler001",
					"level=ERROR msg=\"Failed to PublishCrawlEvent\" crawlerStarterID=starter001 crawlerID=crawler001 err=dummy",
					"level=INFO msg=PublishCrawlEvent crawlerStarterID=starter001 crawlerID=crawler002",
				},
			},
			inputCrawlerStarterSettingID: "starter001",
		},
	}
	for _, tC := range testCases {
		tC.run(t, func(uc *UsecaseImpl) {
			err := uc.StartPipelinePeriodically(context.Background(), tC.inputCrawlerStarterSettingID)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
