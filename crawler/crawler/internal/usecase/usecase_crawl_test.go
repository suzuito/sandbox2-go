package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"go.uber.org/mock/gomock"
)

func TestCrawl(t *testing.T) {
	testCases := []struct {
		desc                  string
		setUp                 func(m *goMocks)
		inputCrawlerID        crawler.CrawlerID
		inputCrawlerInputData crawler.CrawlerInputData
		expectedError         string
	}{
		{
			desc: "success",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				crwl.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any())
				crwl.EXPECT().Parse(gomock.Any(), gomock.Any(), gomock.Any())
				crwl.EXPECT().Publish(gomock.Any(), gomock.Any())
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
			desc: "failed to Fetch",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				crwl.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to Fetch : %+v", gomock.Any()),
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
				crwl.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any())
				crwl.EXPECT().Parse(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
					m.L.EXPECT().Errorf(gomock.Any(), "Failed to Parse : %+v", gomock.Any()),
				)
			},
			inputCrawlerID: "crawler1",
			expectedError:  "dummy error",
		},
		{
			desc: "failed to Publish",
			setUp: func(m *goMocks) {
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				crwl.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any())
				crwl.EXPECT().Parse(gomock.Any(), gomock.Any(), gomock.Any())
				crwl.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(errors.New("dummy error"))
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
			err := u.Crawl(context.Background(), tC.inputCrawlerID, tC.inputCrawlerInputData)
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
				m.Queue.EXPECT().RecieveCrawlEvent(gomock.Any(), gomock.Any()).
					Return(
						crawler.CrawlerID("crawler1"),
						crawler.CrawlerInputData{
							"foo": "bar",
						},
						nil,
					)
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, nil)
				crwl.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any())
				crwl.EXPECT().Parse(gomock.Any(), gomock.Any(), gomock.Any())
				crwl.EXPECT().Publish(gomock.Any(), gomock.Any())
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "Crawl %s", crawler.CrawlerID("crawler1")),
				)
			},
			inputRawBytes: []byte("dummy input"),
		},
		{
			desc: "failed to RecieveCrawlEvent",
			setUp: func(m *goMocks) {
				m.Queue.EXPECT().RecieveCrawlEvent(gomock.Any(), gomock.Any()).
					Return(
						crawler.CrawlerID(""),
						nil,
						fmt.Errorf("dummy error"),
					)
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
				m.Queue.EXPECT().RecieveCrawlEvent(gomock.Any(), gomock.Any()).
					Return(
						crawler.CrawlerID("crawler1"),
						crawler.CrawlerInputData{
							"foo": "bar",
						},
						nil,
					)
				crwl := crawler.NewMockCrawler(m.Controller)
				crwl.EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				crwl.EXPECT().Name().Return("crawler1_name").AnyTimes()
				m.CrawlerFactory.EXPECT().GetCrawler(gomock.Any(), crawler.CrawlerID("crawler1")).
					Return(crwl, errors.New("dummy error"))
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
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
