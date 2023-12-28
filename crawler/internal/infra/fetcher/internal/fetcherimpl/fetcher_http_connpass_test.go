package fetcherimpl

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestFetcherHTTPConnpassID(t *testing.T) {
	assert.Equal(t, (&FetcherHTTPConnpass{}).ID(), crawler.FetcherID("fetcher_http_connpass"))
}

func TestFetcherHTTPConnpassDo(t *testing.T) {
	testCases := []struct {
		desc             string
		mockCli          MockHTTPClientWrapper
		inputTimeNowFunc func() time.Time
		inputQuery       url.Values
		inputDays        int
		expectedError    string
	}{
		{
			desc: "Success",
			mockCli: MockHTTPClientWrapper{
				ExpectedMethod: http.MethodGet,
				ExpectedURL:    mustURL("https://connpass.com/api/v1/event/?hoge=fuga&ymd=20010102&ymd=20010103&ymd=20010104"),
			},
			inputTimeNowFunc: func() time.Time {
				return time.Date(2001, 1, 2, 3, 4, 5, 6, time.UTC)
			},
			inputQuery: url.Values{
				"hoge": []string{"fuga"},
			},
			inputDays: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := FetcherHTTPConnpass{
				Cli:         &tC.mockCli,
				TimeNowFunc: tC.inputTimeNowFunc,
				Query:       tC.inputQuery,
				Days:        tC.inputDays,
			}
			w := bytes.NewBuffer([]byte{})
			logger, _ := clog.NewBytesBufferLogger()
			err := f.Do(context.Background(), logger, w, nil)
			test_helper.AssertError(t, tC.expectedError, err)
			tC.mockCli.Assert(t)
		})
	}
}
