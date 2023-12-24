package factory

import (
	"context"
	"errors"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc                  string
		inputFetcherFactory   FetcherFactory
		inputParserFactory    ParserFactory
		inputPublisherFactory PublisherFactory
		inputDef              crawler.CrawlerDefinition
		expectedError         string
	}{
		{
			desc: "Success",
			inputFetcherFactory: FetcherFactory{
				NewFuncs: []NewFuncFetcher{
					func(_ *crawler.FetcherDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
						return &DummyFetcher{}, nil
					},
				},
			},
			inputParserFactory: ParserFactory{
				NewFuncs: []NewFuncParser{
					func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
						return &DummyParser{}, nil
					},
				},
			},
			inputPublisherFactory: PublisherFactory{
				NewFuncs: []NewFuncPublisher{
					func(_ *crawler.PublisherDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Publisher, error) {
						return &DummyPublisher{}, nil
					},
				},
			},
		},
		{
			desc: "Failed to FetcherFactory.Get",
			inputFetcherFactory: FetcherFactory{
				NewFuncs: []NewFuncFetcher{
					func(_ *crawler.FetcherDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
						return nil, errors.New("dummy")
					},
				},
			},
			expectedError: "dummy",
		},
		{
			desc: "Failed to ParserFactory.Get",
			inputFetcherFactory: FetcherFactory{
				NewFuncs: []NewFuncFetcher{
					func(_ *crawler.FetcherDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
						return &DummyFetcher{}, nil
					},
				},
			},
			inputParserFactory: ParserFactory{
				NewFuncs: []NewFuncParser{
					func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
						return nil, errors.New("dummy")
					},
				},
			},
			expectedError: "dummy",
		},
		{
			desc: "Success",
			inputFetcherFactory: FetcherFactory{
				NewFuncs: []NewFuncFetcher{
					func(_ *crawler.FetcherDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
						return &DummyFetcher{}, nil
					},
				},
			},
			inputParserFactory: ParserFactory{
				NewFuncs: []NewFuncParser{
					func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
						return &DummyParser{}, nil
					},
				},
			},
			inputPublisherFactory: PublisherFactory{
				NewFuncs: []NewFuncPublisher{
					func(_ *crawler.PublisherDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Publisher, error) {
						return nil, errors.New("dummy")
					},
				},
			},
			expectedError: "dummy",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := CrawlerFactory{
				FetcherFactory:   tC.inputFetcherFactory,
				ParserFactory:    tC.inputParserFactory,
				PublisherFactory: tC.inputPublisherFactory,
			}
			_, err := f.Get(context.Background(), &tC.inputDef)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
