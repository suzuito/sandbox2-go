package fetcherimpl

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
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

func TestFetcherHTTPStaticDo(t *testing.T) {
	testCases := []struct {
		desc                    string
		setUp                   func()
		inputReqFunc            func() *http.Request
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
			inputReqFunc: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "https://www.example.com/hoge/", nil)
				return req
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/hoge/ query:map[]]]"`,
			},
		},
		{
			desc: "Error",
			setUp: func() {
				gock.New("https://www.example.com").
					Get("/hoge/").
					ReplyError(errors.New("dummy"))
			},
			inputReqFunc: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "https://www.example.com/hoge/", nil)
				return req
			},
			expectedError: `Get "https://www.example.com/hoge/": dummy`,
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/hoge/ query:map[]]]"`,
			},
		},
		{
			desc: "HTTPError",
			setUp: func() {
				gock.New("https://www.example.com").
					Get("/hoge/").
					Reply(http.StatusNotFound)
			},
			inputReqFunc: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "https://www.example.com/hoge/", nil)
				return req
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedError:           "HTTP error : status=404",
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/hoge/ query:map[]]]"`,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := FetcherHTTPStatic{
				Cli:                http.DefaultClient,
				Req:                tC.inputReqFunc(),
				StatusCodesSuccess: tC.inputStatusCodesSuccess,
			}
			tC.setUp()
			w := bytes.NewBuffer([]byte{})
			logger, logBuffer := newMockLogger()
			err := f.Do(context.Background(), logger, w, nil)
			test_helper.AssertError(t, tC.expectedError, err)
			assertLogString(t, tC.expectedLogLines, logBuffer.String())
		})
	}
}
