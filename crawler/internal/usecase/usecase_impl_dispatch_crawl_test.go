package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestDispatchCrawl(t *testing.T) {
	testCases := []struct {
		utTestCase
		inputCrawlerID crawler.CrawlerID
		expectedError  string
	}{
		{
			utTestCase: utTestCase{
				desc: "Success",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerConfigurationRepository.EXPECT().
						GetDispatchCrawlSetting(gomock.Any()).
						Return(&crawler.DispatchCrawlSetting{
							CrawlFunctionIDMapping: map[crawler.CrawlerID]crawler.CrawlFunctionID{
								"crawler001": "func001",
							},
						}, nil)
					mocks.MockTriggerCrawlerQueue.EXPECT().
						PublishCrawlEvent(gomock.Any(), crawler.CrawlerID("crawler001"), gomock.Any(), crawler.CrawlFunctionID("func001"))
				},
				expectedLogLines: []string{
					"level=INFO msg=DispatchCrawl crawlerID=crawler001",
					"level=INFO msg=PublishCrawlEvent crawlerID=crawler001 crawlFunctionID=func001",
				},
			},
			inputCrawlerID: "crawler001",
		},
		{
			utTestCase: utTestCase{
				desc: "Success (use DefaultFunctionID)",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerConfigurationRepository.EXPECT().
						GetDispatchCrawlSetting(gomock.Any()).
						Return(&crawler.DispatchCrawlSetting{
							CrawlFunctionIDMapping: map[crawler.CrawlerID]crawler.CrawlFunctionID{},
							DefaultCrawlFunctionID: "func002",
						}, nil)
					mocks.MockTriggerCrawlerQueue.EXPECT().
						PublishCrawlEvent(gomock.Any(), crawler.CrawlerID("crawler001"), gomock.Any(), crawler.CrawlFunctionID("func002"))
				},
				expectedLogLines: []string{
					"level=INFO msg=DispatchCrawl crawlerID=crawler001",
					"level=INFO msg=PublishCrawlEvent crawlerID=crawler001 crawlFunctionID=func002",
				},
			},
			inputCrawlerID: "crawler001",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to GetDispatchCrawlSetting",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerConfigurationRepository.EXPECT().
						GetDispatchCrawlSetting(gomock.Any()).
						Return(nil, errors.New("dummy"))
				},
				expectedLogLines: []string{
					"level=INFO msg=DispatchCrawl crawlerID=crawler001",
					`level=ERROR msg="Failed to GetDispatchCrawlSetting" crawlerID=crawler001 err=dummy`,
				},
			},
			inputCrawlerID: "crawler001",
			expectedError:  "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to PublishCrawlEvent",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerConfigurationRepository.EXPECT().
						GetDispatchCrawlSetting(gomock.Any()).
						Return(&crawler.DispatchCrawlSetting{
							CrawlFunctionIDMapping: map[crawler.CrawlerID]crawler.CrawlFunctionID{
								"crawler001": "func001",
							},
						}, nil)
					mocks.MockTriggerCrawlerQueue.EXPECT().
						PublishCrawlEvent(gomock.Any(), crawler.CrawlerID("crawler001"), gomock.Any(), crawler.CrawlFunctionID("func001")).
						Return(errors.New("dummy"))
				},
				expectedLogLines: []string{
					"level=INFO msg=DispatchCrawl crawlerID=crawler001",
					`level=INFO msg=PublishCrawlEvent crawlerID=crawler001 crawlFunctionID=func001`,
					`level=ERROR msg="Failed to PublishCrawlEvent" crawlerID=crawler001 err=dummy`,
				},
			},
			inputCrawlerID: "crawler001",
			expectedError:  "dummy",
		},
	}
	for _, tC := range testCases {
		tC.run(t, func(uc *UsecaseImpl) {
			err := uc.DispatchCrawl(context.Background(), tC.inputCrawlerID, crawler.CrawlerInputData{})
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}

func TestDispatchCrawlOnGCF(t *testing.T) {
	testCases := []struct {
		utTestCase
		expectedError string
	}{
		{
			utTestCase: utTestCase{
				desc: "Success",
				setUp: func(mocks *utMocks) {
					mocks.MockTriggerCrawlerQueue.EXPECT().
						RecieveCrawlEvent(gomock.Any(), gomock.Any()).
						Return(crawler.CrawlerID("crawler001"), crawler.CrawlerInputData{}, nil)
					mocks.MockCrawlerConfigurationRepository.EXPECT().
						GetDispatchCrawlSetting(gomock.Any()).
						Return(&crawler.DispatchCrawlSetting{
							CrawlFunctionIDMapping: map[crawler.CrawlerID]crawler.CrawlFunctionID{
								"crawler001": "func001",
							},
						}, nil)
					mocks.MockTriggerCrawlerQueue.EXPECT().
						PublishCrawlEvent(gomock.Any(), crawler.CrawlerID("crawler001"), gomock.Any(), crawler.CrawlFunctionID("func001"))
				},
				expectedLogLines: []string{
					"level=INFO msg=DispatchCrawl crawlerID=crawler001",
					"level=INFO msg=PublishCrawlEvent crawlerID=crawler001 crawlFunctionID=func001",
				},
			},
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to RecieveCrawlEvent",
				setUp: func(mocks *utMocks) {
					mocks.MockTriggerCrawlerQueue.EXPECT().
						RecieveCrawlEvent(gomock.Any(), gomock.Any()).
						Return(crawler.CrawlerID(""), crawler.CrawlerInputData{}, errors.New("dummy"))
				},
				expectedLogLines: []string{
					`level=ERROR msg="Failed to RecieveCrawlEvent" err=dummy`,
				},
			},
			expectedError: "dummy",
		},
	}
	for _, tC := range testCases {
		tC.run(t, func(uc *UsecaseImpl) {
			err := uc.DispatchCrawlOnGCF(context.Background(), []byte{})
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
