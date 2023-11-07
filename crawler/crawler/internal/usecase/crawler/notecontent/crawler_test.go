package notecontent

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
	"go.uber.org/mock/gomock"
)

func TestNewCrawler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := NewCrawler(
		crawler.CrawlerID("dummy_id"),
		repository.NewMockRepository(ctrl),
		fetcher.NewMockFetcherHTTP(ctrl),
		func(article *note.TimeSeriesDataNoteArticle) bool { return true },
	)
	assert.Equal(t, crawler.CrawlerID("dummy_id"), c.ID())
	assert.Equal(t, "dummy_id", c.Name())
}

func TestFetch(t *testing.T) {
	testCases := []struct {
		desc        string
		inputInput  crawler.CrawlerInputData
		setUp       func(mockFetcher *fetcher.MockFetcherHTTP)
		expectedErr string
	}{
		{
			desc: "Success to fetch",
			inputInput: crawler.CrawlerInputData{
				"URL": "https://note.com/kotton_yururun/n/n4cdb54623e48",
			},
			setUp: func(mockFetcher *fetcher.MockFetcherHTTP) {
				mockFetcher.EXPECT().DoRequest(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				)
			},
		},
		{
			desc:       "URL not found in input",
			inputInput: crawler.CrawlerInputData{},
			setUp: func(mockFetcher *fetcher.MockFetcherHTTP) {
			},
			expectedErr: `input\["URL"\] not found in input`,
		},
		{
			desc: "Invalid URL in input",
			inputInput: crawler.CrawlerInputData{
				"URL": ":",
			},
			setUp: func(mockFetcher *fetcher.MockFetcherHTTP) {
			},
			expectedErr: `parse ":": missing protocol scheme`,
		},
		{
			desc: "Failed to fetch",
			inputInput: crawler.CrawlerInputData{
				"URL": "https://note.com/kotton_yururun/n/n4cdb54623e48",
			},
			setUp: func(mockFetcher *fetcher.MockFetcherHTTP) {
				mockFetcher.EXPECT().DoRequest(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(errors.New("dummy error"))
			},
			expectedErr: "dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockFetcher := fetcher.NewMockFetcherHTTP(ctrl)
			tC.setUp(mockFetcher)
			c := NewCrawler(
				crawler.CrawlerID("dummy_id"),
				repository.NewMockRepository(ctrl),
				mockFetcher,
				func(article *note.TimeSeriesDataNoteArticle) bool { return true },
			)
			err := c.Fetch(context.Background(), bytes.NewBufferString(""), tC.inputInput)
			test_helper.AssertError(t, tC.expectedErr, err)
		})
	}
}
