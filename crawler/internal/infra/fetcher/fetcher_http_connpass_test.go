package fetcher

import (
	"net/url"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestNewFetcherHTTPConnpass(t *testing.T) {
	testCases := []struct {
		desc          string
		inputDef      crawler.FetcherDefinition
		expectedError string
	}{
		{
			desc: "Success",
			inputDef: crawler.FetcherDefinition{
				ID: "fetcher_http_connpass",
				Argument: argument.ArgumentDefinition{
					"Days": 3,
					"Query": url.Values{
						"a": []string{"b"},
					},
				},
			},
		},
		{
			desc: "Not matched id",
			inputDef: crawler.FetcherDefinition{
				ID: "foo",
			},
			expectedError: `NoMatchedFetcherID`,
		},
		{
			desc: "No Days in argument",
			inputDef: crawler.FetcherDefinition{
				ID: "fetcher_http_connpass",
				Argument: argument.ArgumentDefinition{
					"Query": url.Values{
						"a": []string{"b"},
					},
				},
			},
			expectedError: `Key 'Days' is not found in AgumentDefinition`,
		},
		{
			desc: "No Query in argument",
			inputDef: crawler.FetcherDefinition{
				ID: "fetcher_http_connpass",
				Argument: argument.ArgumentDefinition{
					"Days": 3,
				},
			},
			expectedError: `Key 'Query' is not found in AgumentDefinition`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := NewFetcherHTTPConnpass(
				&tC.inputDef,
				&factorysetting.CrawlerFactorySetting{},
			)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
