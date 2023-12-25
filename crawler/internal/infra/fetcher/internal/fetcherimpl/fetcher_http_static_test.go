package fetcherimpl

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/suzuito/sandbox2-go/common/test_helper"
)

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
