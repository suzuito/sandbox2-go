package notecontent

import (
	"bytes"
	"context"
	"errors"
	"io"
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

func TestParse(t *testing.T) {
	testCases := []struct {
		desc                     string
		inputR                   io.Reader
		inputFilter              func(article *note.TimeSeriesDataNoteArticle) bool
		expectedNoteArticleTitle string
		expectedErr              string
	}{
		{
			desc: "Success",
			inputFilter: func(article *note.TimeSeriesDataNoteArticle) bool {
				return true
			},
			inputR: bytes.NewBufferString(`
<!doctype html>
<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
<head>
    <title>title1</title>
    <link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
    <meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
    <meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
</head>
<body>
    <div class="o-noteContentHeader__info">
        <div class="o-noteContentHeader__name">
            <time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
        </div>
        <ul id="tagListBody" class="m-tagList__body">
            <li tabindex="-1" class="m-tagList__item" style="display:;">
                <div class="a-tag__label">
                    #デザイン
                    <!---->
                </div>
            </li>
            <li tabindex="-1" class="m-tagList__item" style="display:;">
                <div class="a-tag__label">
                    #デザイナー
                    <!---->
                </div>
            </li>
            <li tabindex="-1" class="m-tagList__item" style="display:;">
                <div class="a-tag__label">
                    #ナレッジワーク
                </div>
            </li>
        </ul>
        <div class="p-article__content" data-v-c5502208>
            This is content
        </div>
</body>
</html>
			`),
			expectedNoteArticleTitle: "title1",
		},
		{
			desc: "Failed to parse",
			inputFilter: func(article *note.TimeSeriesDataNoteArticle) bool {
				return true
			},
			inputR:      bytes.NewBufferString(`hoge`),
			expectedErr: `Cannot find link\[rel=canonical\] tag`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			c := NewCrawler(
				crawler.CrawlerID("dummy_id"),
				repository.NewMockRepository(ctrl),
				fetcher.NewMockFetcherHTTP(ctrl),
				tC.inputFilter,
			)
			actuals, err := c.Parse(context.Background(), tC.inputR, nil)
			test_helper.AssertError(t, tC.expectedErr, err)
			if err == nil {
				actual := actuals[0].(*note.TimeSeriesDataNoteArticle)
				assert.Equal(t, tC.expectedNoteArticleTitle, actual.Title)
			}
		})
	}
}

func TestPublish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockRepository(ctrl)
	repo.EXPECT().SetTimeSeriesData(gomock.Any(), crawler.CrawlerID("dummy_id"), gomock.Any())
	c := NewCrawler(
		crawler.CrawlerID("dummy_id"),
		repo,
		fetcher.NewMockFetcherHTTP(ctrl),
		func(article *note.TimeSeriesDataNoteArticle) bool { return true },
	)
	c.Publish(context.Background(), nil, &note.TimeSeriesDataNoteArticle{})
}
