package fetcher

import (
	"net/http"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestNewFetcherHTTPStatic(t *testing.T) {
	testCases := []struct {
		desc          string
		inputDef      crawler.FetcherDefinition
		expectedError string
	}{
		{
			desc: "Success",
			inputDef: crawler.FetcherDefinition{
				ID: "fetcher_http_static",
				Argument: argument.ArgumentDefinition{
					"URL":                "https://www.example.com/hoge",
					"Method":             "GET",
					"StatusCodesSuccess": []int{http.StatusOK},
				},
			},
		},
		{
			desc: "URL is not found in argument",
			inputDef: crawler.FetcherDefinition{
				ID: "foo",
			},
			expectedError: "NoMatchedFetcherID",
		},
		{
			desc: "URL is not found in argument",
			inputDef: crawler.FetcherDefinition{
				ID:       "fetcher_http_static",
				Argument: argument.ArgumentDefinition{},
			},
			expectedError: `Key 'URL' is not found in AgumentDefinition`,
		},
		{
			desc: "Method is not found in argument",
			inputDef: crawler.FetcherDefinition{
				ID: "fetcher_http_static",
				Argument: argument.ArgumentDefinition{
					"URL": "https://www.example.com/hoge",
				},
			},
			expectedError: `Key 'Method' is not found in AgumentDefinition`,
		},
		{
			desc: "StatusCodesSuccess is not found in argument",
			inputDef: crawler.FetcherDefinition{
				ID: "fetcher_http_static",
				Argument: argument.ArgumentDefinition{
					"URL":    "https://www.example.com/hoge",
					"Method": "GET",
				},
			},
			expectedError: `Key 'StatusCodesSuccess' is not found in AgumentDefinition`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := NewFetcherHTTPStatic(&tC.inputDef, &factorysetting.CrawlerFactorySetting{})
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
