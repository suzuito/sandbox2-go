package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/entity/crawler"
	"go.uber.org/mock/gomock"
)

func TestStartPipelinePeriodically(t *testing.T) {
	testCases := []struct {
		desc          string
		setUp         func(m *goMocks)
		expectedError string
	}{
		{
			desc: "0 crawlers",
			setUp: func(m *goMocks) {
				crawlers := []crawler.Crawler{}
				m.CrawlerFactory.EXPECT().
					GetCrawlers(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).
					Return(crawlers)
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "StartPipelinePeriodically"),
				)
			},
		},
		{
			desc: "3 crawlers",
			setUp: func(m *goMocks) {
				mockCrawlers := []*crawler.MockCrawler{
					crawler.NewMockCrawler(m.Controller),
					crawler.NewMockCrawler(m.Controller),
					crawler.NewMockCrawler(m.Controller),
				}
				mockCrawlers[0].EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				mockCrawlers[0].EXPECT().Name().Return("crawler1_name").AnyTimes()
				mockCrawlers[1].EXPECT().ID().Return(crawler.CrawlerID("crawler2")).AnyTimes()
				mockCrawlers[1].EXPECT().Name().Return("crawler2_name").AnyTimes()
				mockCrawlers[2].EXPECT().ID().Return(crawler.CrawlerID("crawler3")).AnyTimes()
				mockCrawlers[2].EXPECT().Name().Return("crawler3_name").AnyTimes()
				crawlers := []crawler.Crawler{}
				for _, mc := range mockCrawlers {
					crawlers = append(crawlers, mc)
				}
				m.CrawlerFactory.EXPECT().
					GetCrawlers(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).
					Return(crawlers)
				gomock.InOrder(
					m.Queue.EXPECT().
						PublishCrawlEvent(
							gomock.Any(),
							crawler.CrawlerID("crawler1"),
						),
					m.Queue.EXPECT().
						PublishCrawlEvent(
							gomock.Any(),
							crawler.CrawlerID("crawler2"),
						),
					m.Queue.EXPECT().
						PublishCrawlEvent(
							gomock.Any(),
							crawler.CrawlerID("crawler3"),
						),
				)
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "StartPipelinePeriodically"),
					m.L.EXPECT().Infof(gomock.Any(), "Start %s (%s)", crawler.CrawlerID("crawler1"), "crawler1_name"),
					m.L.EXPECT().Infof(gomock.Any(), "Start %s (%s)", crawler.CrawlerID("crawler2"), "crawler2_name"),
					m.L.EXPECT().Infof(gomock.Any(), "Start %s (%s)", crawler.CrawlerID("crawler3"), "crawler3_name"),
				)
			},
		},
		{
			desc: "3 crawlers and 1 pub error",
			setUp: func(m *goMocks) {
				mockCrawlers := []*crawler.MockCrawler{
					crawler.NewMockCrawler(m.Controller),
					crawler.NewMockCrawler(m.Controller),
					crawler.NewMockCrawler(m.Controller),
				}
				mockCrawlers[0].EXPECT().ID().Return(crawler.CrawlerID("crawler1")).AnyTimes()
				mockCrawlers[0].EXPECT().Name().Return("crawler1_name").AnyTimes()
				mockCrawlers[1].EXPECT().ID().Return(crawler.CrawlerID("crawler2")).AnyTimes()
				mockCrawlers[1].EXPECT().Name().Return("crawler2_name").AnyTimes()
				mockCrawlers[2].EXPECT().ID().Return(crawler.CrawlerID("crawler3")).AnyTimes()
				mockCrawlers[2].EXPECT().Name().Return("crawler3_name").AnyTimes()
				crawlers := []crawler.Crawler{}
				for _, mc := range mockCrawlers {
					crawlers = append(crawlers, mc)
				}
				m.CrawlerFactory.EXPECT().
					GetCrawlers(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).
					Return(crawlers)
				gomock.InOrder(
					m.Queue.EXPECT().
						PublishCrawlEvent(
							gomock.Any(),
							crawler.CrawlerID("crawler1"),
						),
					m.Queue.EXPECT().
						PublishCrawlEvent(
							gomock.Any(),
							crawler.CrawlerID("crawler2"),
						).
						Return(fmt.Errorf("dummy error")),
					m.Queue.EXPECT().
						PublishCrawlEvent(
							gomock.Any(),
							crawler.CrawlerID("crawler3"),
						),
				)
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "StartPipelinePeriodically"),
					m.L.EXPECT().Infof(gomock.Any(), "Start %s (%s)", crawler.CrawlerID("crawler1"), "crawler1_name"),
					m.L.EXPECT().Infof(gomock.Any(), "Start %s (%s)", crawler.CrawlerID("crawler2"), "crawler2_name"),
					m.L.EXPECT().Errorf(gomock.Any(), "PublishCrawlEvent of '%s' is failed : %+v", crawler.CrawlerID("crawler2"), gomock.Any()),
					m.L.EXPECT().Infof(gomock.Any(), "Start %s (%s)", crawler.CrawlerID("crawler3"), "crawler3_name"),
				)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mocks := newMocks(t)
			defer mocks.Finish()
			tC.setUp(mocks)
			u := mocks.NewUsecase()
			err := u.StartPipelinePeriodically(ctx)
			test_helper.AssertErrorAs(t, tC.expectedError, err)
		})
	}
}
