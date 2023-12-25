package fetcher

import (
	"net/http"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestNewFetcherHTTP(t *testing.T) {
	testCases := []struct {
		desc          string
		inputDef      crawler.FetcherDefinition
		expectedError string
	}{
		{
			desc: "Success",
			inputDef: crawler.FetcherDefinition{
				ID: "fetcher_http",
				Argument: argument.ArgumentDefinition{
					"StatusCodesSuccess": []int{http.StatusOK},
				},
			},
		},
		{
			desc: "Invalid ID",
			inputDef: crawler.FetcherDefinition{
				ID: "invalid_id",
				Argument: argument.ArgumentDefinition{
					"StatusCodesSuccess": []int{http.StatusOK},
				},
			},
			expectedError: "NoMatchedFetcherID",
		},
		{
			desc: "Missing StatusCodesSuccess",
			inputDef: crawler.FetcherDefinition{
				ID:       "fetcher_http",
				Argument: argument.ArgumentDefinition{
					// "StatusCodesSuccess" key is missing
				},
			},
			expectedError: "Key 'StatusCodesSuccess' is not found in AgumentDefinition",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := NewFetcherHTTP(&tC.inputDef, &factorysetting.CrawlerFactorySetting{})
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
