package fetcherimpl

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestNewFetcherHTTP(t *testing.T) {
	testCases := []struct {
		desc          string
		inputDef      crawler.FetcherDefinition
		inputArgs     factory.NewFuncFetcherArgument
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
			inputArgs: factory.NewFuncFetcherArgument{},
		},
		{
			desc: "Invalid ID",
			inputDef: crawler.FetcherDefinition{
				ID: "invalid_id",
				Argument: argument.ArgumentDefinition{
					"StatusCodesSuccess": []int{http.StatusOK},
				},
			},
			inputArgs:     factory.NewFuncFetcherArgument{},
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
			inputArgs:     factory.NewFuncFetcherArgument{},
			expectedError: "Key 'StatusCodesSuccess' is not found in AgumentDefinition",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := NewFetcherHTTP(&tC.inputDef, &tC.inputArgs)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}

func TestFetcherHTTPDo(t *testing.T) {
	testCases := []struct {
		desc                    string
		setUp                   func()
		inputCrawlerInputData   crawler.CrawlerInputData
		inputStatusCodesSuccess []int
		expectedLogLines        []string
		expectedError           string
	}{
		{
			desc: "Success",
			setUp: func() {
				gock.New("https://www.example.com").
					Get("/hoge/").
					Reply(http.StatusOK)
			},
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL":    "https://www.example.com/hoge/",
				"Method": "GET",
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/hoge/ query:map[]]]"`,
			},
		},
		{
			desc: "Success (omit Method)",
			setUp: func() {
				gock.New("https://www.example.com").
					Get("/hoge/").
					Reply(http.StatusOK)
			},
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL": "https://www.example.com/hoge/",
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/hoge/ query:map[]]]"`,
			},
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
		{
			desc: "Error - HTTP client failed",
			setUp: func() {
				gock.New("https://www.example.com").
					Get("/hoge/").
					ReplyError(errors.New("dummy"))
			},
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL": "https://www.example.com/hoge/",
			},
			expectedError: "Get \"https://www.example.com/hoge/\": dummy",
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/hoge/ query:map[]]]"`,
			},
		},
		{
			desc: "Error - HTTP request failed",
			setUp: func() {
				gock.New("https://www.example.com").
					Get("/error/").
					Reply(http.StatusNotFound)
			},
			inputCrawlerInputData: crawler.CrawlerInputData{
				"URL":    "https://www.example.com/error/",
				"Method": "GET",
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedError:           "HTTP error : status=404",
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/error/ query:map[]]]"`,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := FetcherHTTP{
				Cli:                http.DefaultClient,
				StatusCodesSuccess: tC.inputStatusCodesSuccess,
			}
			if tC.setUp != nil {
				tC.setUp()
			}
			w := bytes.NewBuffer([]byte{})
			logger, logBuffer := newMockLogger()
			err := f.Do(context.Background(), logger, w, tC.inputCrawlerInputData)
			test_helper.AssertError(t, tC.expectedError, err)
			assertLogString(t, tC.expectedLogLines, logBuffer.String())
		})
	}
}
