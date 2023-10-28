package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"go.uber.org/mock/gomock"
)

func TestCrawl(t *testing.T) {
	testCases := []struct {
		desc           string
		setUp          func(m *goMocks)
		inputCrawlerID crawler.CrawlerID
		expectedError  string
	}{
		{
			desc: "success",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				fetcher := crawler.NewMockFetcher(m.Controller)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(fetcher, nil)
				fetcher.EXPECT().Fetch(gomock.Any(), gomock.Any())
				parser := crawler.NewMockParser(m.Controller)
				crwl.EXPECT().NewParser(gomock.Any()).Return(parser, nil)
				parser.EXPECT().Parse(gomock.Any(), gomock.Any())
				publisher := crawler.NewMockPublisher(m.Controller)
				crwl.EXPECT().NewPublisher(gomock.Any()).Return(publisher, nil)
				publisher.EXPECT().Publish(gomock.Any(), gomock.Any())
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
				)
			},
			inputCrawlerID: "crawler1",
		},
		{
			desc: "failed to GetCrawler",
			setUp: func(m *goMocks) {
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(nil, fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to GetCrawler : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
		{
			desc: "failed to NewFetcher",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(nil, fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to NewFetcher : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
		{
			desc: "failed to Fetch",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				fetcher := crawler.NewMockFetcher(m.Controller)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(fetcher, nil)
				fetcher.EXPECT().Fetch(gomock.Any(), gomock.Any()).Return(fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to Fetch : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
		{
			desc: "failed to NewParser",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				fetcher := crawler.NewMockFetcher(m.Controller)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(fetcher, nil)
				fetcher.EXPECT().Fetch(gomock.Any(), gomock.Any())
				parser := crawler.NewMockParser(m.Controller)
				crwl.EXPECT().NewParser(gomock.Any()).Return(parser, nil).Return(nil, fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to NewParser : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
		{
			desc: "failed to Parse",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				fetcher := crawler.NewMockFetcher(m.Controller)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(fetcher, nil)
				fetcher.EXPECT().Fetch(gomock.Any(), gomock.Any())
				parser := crawler.NewMockParser(m.Controller)
				crwl.EXPECT().NewParser(gomock.Any()).Return(parser, nil)
				parser.EXPECT().Parse(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to Parse : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
		{
			desc: "failed to NewPublisher",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				fetcher := crawler.NewMockFetcher(m.Controller)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(fetcher, nil)
				fetcher.EXPECT().Fetch(gomock.Any(), gomock.Any())
				parser := crawler.NewMockParser(m.Controller)
				crwl.EXPECT().NewParser(gomock.Any()).Return(parser, nil)
				parser.EXPECT().Parse(gomock.Any(), gomock.Any())
				publisher := crawler.NewMockPublisher(m.Controller)
				crwl.EXPECT().NewPublisher(gomock.Any()).Return(publisher, nil).Return(nil, fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to NewPublisher : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
		{
			desc: "success",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				fetcher := crawler.NewMockFetcher(m.Controller)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(fetcher, nil)
				fetcher.EXPECT().Fetch(gomock.Any(), gomock.Any())
				parser := crawler.NewMockParser(m.Controller)
				crwl.EXPECT().NewParser(gomock.Any()).Return(parser, nil)
				parser.EXPECT().Parse(gomock.Any(), gomock.Any())
				publisher := crawler.NewMockPublisher(m.Controller)
				crwl.EXPECT().NewPublisher(gomock.Any()).Return(publisher, nil)
				publisher.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to Publish : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := newMocks(t)
			defer m.Finish()
			tC.setUp(m)
			u := m.NewUsecase()
			err := u.Crawl(context.Background(), tC.inputCrawlerID)
			if err == nil {
			}
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}

func TestCrawlOnGCF(t *testing.T) {
	testCases := []struct {
		desc          string
		setUp         func(m *goMocks)
		inputRawBytes []byte
		expectedError string
	}{
		{
			desc: "success",
			setUp: func(m *goMocks) {
				m.Queue.EXPECT().RecieveCrawlEvent(gomock.Any(), gomock.Any()).Return(crawler.CrawlerID("crawler1"), nil)
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				fetcher := crawler.NewMockFetcher(m.Controller)
				crwl.EXPECT().NewFetcher(gomock.Any()).Return(fetcher, nil)
				fetcher.EXPECT().Fetch(gomock.Any(), gomock.Any())
				parser := crawler.NewMockParser(m.Controller)
				crwl.EXPECT().NewParser(gomock.Any()).Return(parser, nil)
				parser.EXPECT().Parse(gomock.Any(), gomock.Any())
				publisher := crawler.NewMockPublisher(m.Controller)
				crwl.EXPECT().NewPublisher(gomock.Any()).Return(publisher, nil)
				publisher.EXPECT().Publish(gomock.Any(), gomock.Any())
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
				)
			},
			inputRawBytes: []byte("dummy input"),
		},
		{
			desc: "failed to RecieveCrawlEvent",
			setUp: func(m *goMocks) {
				m.Queue.EXPECT().RecieveCrawlEvent(gomock.Any(), gomock.Any()).Return(crawler.CrawlerID(""), fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to RecieveCrawlEvent : %+v", gomock.Any()),
				)
			},
			inputRawBytes: []byte("dummy input"),
			expectedError: "dummy error",
		},
		{
			desc: "failed to GetCrawler",
			setUp: func(m *goMocks) {
				m.Queue.EXPECT().RecieveCrawlEvent(gomock.Any(), gomock.Any()).Return(crawler.CrawlerID("crawler1"), nil)
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(nil, fmt.Errorf("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to GetCrawler : %+v", gomock.Any()),
				)
			},
			inputRawBytes: []byte("dummy input"),
			expectedError: "dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := newMocks(t)
			defer m.Finish()
			tC.setUp(m)
			u := m.NewUsecase()
			err := u.CrawlOnGCF(context.Background(), tC.inputRawBytes)
			if err == nil {
			}
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
