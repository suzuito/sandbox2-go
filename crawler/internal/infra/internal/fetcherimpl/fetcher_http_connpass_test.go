package fetcherimpl

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestNewFetcherHTTPConnpass(t *testing.T) {
	testCases := []struct {
		desc          string
		inputDef      crawler.FetcherDefinition
		inputArgs     factory.NewFuncFetcherArgument
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
				&tC.inputArgs,
			)
			test_helper.AssertError(t, tC.expectedError, err)
			if err != nil {
				return
			}
		})
	}
}

func TestFetcherHTTPConnpassDo(t *testing.T) {
	testCases := []struct {
		desc             string
		setUp            func()
		inputTimeNowFunc func() time.Time
		inputQuery       url.Values
		inputDays        int
		expectedError    string
	}{
		{
			desc: "Success",
			inputTimeNowFunc: func() time.Time {
				return time.Date(2001, 1, 2, 3, 4, 5, 6, time.UTC)
			},
			inputQuery: url.Values{
				"hoge": []string{"fuga"},
			},
			inputDays: 3,
			setUp: func() {
				gock.New("https://connpass.com").
					Get("/api/v1/event/").
					MatchParam("hoge", "fuga").
					MatchParam("ymd", "20010102").
					MatchParam("ymd", "20010103").
					MatchParam("ymd", "20010104").
					Reply(http.StatusOK)
			},
		},
		{
			desc: "Unknown error when http request Do func",
			inputTimeNowFunc: func() time.Time {
				return time.Date(2001, 1, 2, 3, 4, 5, 6, time.UTC)
			},
			inputQuery: url.Values{
				"hoge": []string{"fuga"},
			},
			inputDays: 3,
			setUp: func() {
				gock.New("https://connpass.com").
					Get("/api/v1/event/").
					MatchParam("hoge", "fuga").
					MatchParam("ymd", "20010102").
					MatchParam("ymd", "20010103").
					MatchParam("ymd", "20010104").
					ReplyError(errors.New("dummy"))
			},
			expectedError: `dummy`,
		},
		{
			desc: "HTTP error (not 200)",
			inputTimeNowFunc: func() time.Time {
				return time.Date(2001, 1, 2, 3, 4, 5, 6, time.UTC)
			},
			inputQuery: url.Values{
				"hoge": []string{"fuga"},
			},
			inputDays: 3,
			setUp: func() {
				gock.New("https://connpass.com").
					Get("/api/v1/event/").
					MatchParam("hoge", "fuga").
					MatchParam("ymd", "20010102").
					MatchParam("ymd", "20010103").
					MatchParam("ymd", "20010104").
					Reply(http.StatusNotFound)
			},
			expectedError: `HTTP error : status=404`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer gock.Off()
			f := FetcherHTTPConnpass{
				Cli:         http.DefaultClient,
				TimeNowFunc: tC.inputTimeNowFunc,
				Query:       tC.inputQuery,
				Days:        tC.inputDays,
			}
			tC.setUp()
			w := bytes.NewBuffer([]byte{})
			err := f.Do(context.Background(), w, nil)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
