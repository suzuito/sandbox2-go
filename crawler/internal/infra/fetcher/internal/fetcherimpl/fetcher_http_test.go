package fetcherimpl

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestFetcherHTTPDo(t *testing.T) {
	testCases := []struct {
		desc                    string
		mockCli                 *MockHTTPClientWrapper
		inputCrawlerInputData   crawler.CrawlerInputData
		inputStatusCodesSuccess []int
		expectedError           string
	}{
		{
			desc: "Success",
			mockCli: &MockHTTPClientWrapper{
				ExpectedMethod: http.MethodGet,
				ExpectedURL:    mustURL("https://www.example.com/hoge/"),
			},
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL":    "https://www.example.com/hoge/",
				"Method": "GET",
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
		},
		{
			desc: "Success (omit Method)",
			mockCli: &MockHTTPClientWrapper{
				ExpectedMethod: http.MethodGet,
				ExpectedURL:    mustURL("https://www.example.com/hoge/"),
			},
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL": "https://www.example.com/hoge/",
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
		},
		{
			desc: "Error - URL not found",
			inputCrawlerInputData: crawler.CrawlerInputData{
				// "URL" key is missing
				"Method": "GET",
			},
			expectedError: `input\["URL"\] not found in input`,
		},
		{
			desc: "Error - Invalid method type",
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL":    "https://www.example.com/hoge/",
				"Method": 123, // Invalid method type (not a string)
			},
			expectedError: `input\["Method"\] must be string in input`,
		},
		{
			desc: "Error - Invalid URL",
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL":    ":::", // Invalid URL
				"Method": "GET",
			},
			expectedError: "parse \":::\": missing protocol scheme",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := FetcherHTTP{
				Cli:                tC.mockCli,
				StatusCodesSuccess: tC.inputStatusCodesSuccess,
			}
			w := bytes.NewBuffer([]byte{})
			logger, _ := newMockLogger()
			err := f.Do(context.Background(), logger, w, tC.inputCrawlerInputData)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.mockCli != nil {
				tC.mockCli.Assert(t)
			}
		})
	}
}
