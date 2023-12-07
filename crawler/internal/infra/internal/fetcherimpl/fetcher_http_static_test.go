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

func TestNewFetcherHTTPStatic(t *testing.T) {
	testCases := []struct {
		desc          string
		inputDef      crawler.FetcherDefinition
		inputArgs     factory.NewFuncFetcherArgument
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
			inputArgs: factory.NewFuncFetcherArgument{},
		},
		{
			desc: "URL is not found in argument",
			inputDef: crawler.FetcherDefinition{
				ID: "foo",
			},
			inputArgs:     factory.NewFuncFetcherArgument{},
			expectedError: "NoMatchedFetcherID",
		},
		{
			desc: "URL is not found in argument",
			inputDef: crawler.FetcherDefinition{
				ID:       "fetcher_http_static",
				Argument: argument.ArgumentDefinition{},
			},
			inputArgs:     factory.NewFuncFetcherArgument{},
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
			inputArgs:     factory.NewFuncFetcherArgument{},
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
			inputArgs:     factory.NewFuncFetcherArgument{},
			expectedError: `Key 'StatusCodesSuccess' is not found in AgumentDefinition`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := NewFetcherHTTPStatic(&tC.inputDef, &tC.inputArgs)
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
			err := f.Do(context.Background(), w, nil)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
