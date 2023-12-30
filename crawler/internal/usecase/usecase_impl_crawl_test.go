package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"go.uber.org/mock/gomock"
)

func TestCrawl(t *testing.T) {
	testCases := []struct {
		utTestCase
		inputCrawlerID        crawler.CrawlerID
		inputCrawlerInputData crawler.CrawlerInputData
		expectedError         string
	}{
		{
			utTestCase: utTestCase{
				desc: "Success",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerDefinition(
							gomock.Any(),
							crawler.CrawlerID("crawler001"),
						).
						Return(&crawler.CrawlerDefinition{}, nil)
					mocks.MockFetcher.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					)
					mocks.MockParser.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).Return(
						[]timeseriesdata.TimeSeriesData{
							&timeseriesdata.TimeSeriesDataEvent{},
						},
						nil,
					)
					mocks.MockPublisher.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					)
					crwl := crawler.Crawler{
						Fetcher:   mocks.MockFetcher,
						Parser:    mocks.MockParser,
						Publisher: mocks.MockPublisher,
					}
					mocks.MockCrawlerFactory.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(&crwl, nil)
				},
				expectedLogLines: []string{
					"level=INFO msg=Crawl crawlerID=crawler001",
				},
			},
			inputCrawlerID: crawler.CrawlerID("crawler001"),
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to t.CrawlerRepository.GetCrawlerDefinition",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerDefinition(
							gomock.Any(),
							crawler.CrawlerID("crawler001"),
						).
						Return(nil, errors.New("dummy"))
				},
				expectedLogLines: []string{
					"level=INFO msg=Crawl crawlerID=crawler001",
					`level=ERROR msg="Failed to GetCrawler" crawlerID=crawler001 err=dummy`,
				},
			},
			inputCrawlerID: crawler.CrawlerID("crawler001"),
			expectedError:  "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to t.CrawlerFactory.Get",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerDefinition(
							gomock.Any(),
							crawler.CrawlerID("crawler001"),
						).
						Return(&crawler.CrawlerDefinition{}, nil)
					mocks.MockCrawlerFactory.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(nil, errors.New("dummy"))
				},
				expectedLogLines: []string{
					"level=INFO msg=Crawl crawlerID=crawler001",
					`level=ERROR msg="Failed to CrawlerFactory.Get" crawlerID=crawler001 err=dummy`,
				},
			},
			inputCrawlerID: crawler.CrawlerID("crawler001"),
			expectedError:  "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to Fetcher.Do",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerDefinition(
							gomock.Any(),
							crawler.CrawlerID("crawler001"),
						).
						Return(&crawler.CrawlerDefinition{}, nil)
					mocks.MockFetcher.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).Return(errors.New("dummy"))
					crwl := crawler.Crawler{
						Fetcher: mocks.MockFetcher,
					}
					mocks.MockCrawlerFactory.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(&crwl, nil)
				},
				expectedLogLines: []string{
					"level=INFO msg=Crawl crawlerID=crawler001",
					`level=ERROR msg="Failed to Fetch" crawlerID=crawler001 err=dummy`,
				},
			},
			inputCrawlerID: crawler.CrawlerID("crawler001"),
			expectedError:  "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to Parser.Do",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerDefinition(
							gomock.Any(),
							crawler.CrawlerID("crawler001"),
						).
						Return(&crawler.CrawlerDefinition{}, nil)
					mocks.MockFetcher.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					)
					mocks.MockParser.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).Return(nil, errors.New("dummy"))
					crwl := crawler.Crawler{
						Fetcher: mocks.MockFetcher,
						Parser:  mocks.MockParser,
					}
					mocks.MockCrawlerFactory.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(&crwl, nil)
				},
				expectedLogLines: []string{
					"level=INFO msg=Crawl crawlerID=crawler001",
					`level=ERROR msg="Failed to Parse" crawlerID=crawler001 err=dummy`,
				},
			},
			inputCrawlerID: crawler.CrawlerID("crawler001"),
			expectedError:  "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to Publisher.Do",
				setUp: func(mocks *utMocks) {
					mocks.MockCrawlerRepository.EXPECT().
						GetCrawlerDefinition(
							gomock.Any(),
							crawler.CrawlerID("crawler001"),
						).
						Return(&crawler.CrawlerDefinition{}, nil)
					mocks.MockFetcher.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					)
					mocks.MockParser.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).Return(
						[]timeseriesdata.TimeSeriesData{
							&timeseriesdata.TimeSeriesDataEvent{},
						},
						nil,
					)
					mocks.MockPublisher.EXPECT().Do(
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).Return(errors.New("dummy"))
					crwl := crawler.Crawler{
						Fetcher:   mocks.MockFetcher,
						Parser:    mocks.MockParser,
						Publisher: mocks.MockPublisher,
					}
					mocks.MockCrawlerFactory.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(&crwl, nil)
				},
				expectedLogLines: []string{
					"level=INFO msg=Crawl crawlerID=crawler001",
					`level=ERROR msg="Failed to Publish" crawlerID=crawler001 err=dummy`,
				},
			},
			inputCrawlerID: crawler.CrawlerID("crawler001"),
			expectedError:  "dummy",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.run(t, func(uc *UsecaseImpl) {
				err := uc.Crawl(context.Background(), tC.inputCrawlerID, tC.inputCrawlerInputData)
				test_helper.AssertError(t, tC.expectedError, err)
			})
		})
	}
}
