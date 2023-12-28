package httpclientwrapperimpl

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/httpclientcache"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"go.uber.org/mock/gomock"
)

func Test(t *testing.T) {
	inputReq, _ := http.NewRequest("GET", "https://www.example.com/", nil)
	testCases := []struct {
		desc                    string
		setUp                   func(mockCache *httpclientcache.MockClient)
		inputUseCache           bool
		inputStatusCodesSuccess []int
		expectedLogLines        []string
		expectedError           string
	}{
		{
			desc: "Success(cache=false)",
			setUp: func(mockCache *httpclientcache.MockClient) {
				gock.New("https://www.example.com").
					Get("/").
					Reply(http.StatusOK).
					BodyString("Hello, world!")
			},
			inputUseCache:           false,
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/ query:map[]]]"`,
			},
		},
		{
			desc: "Success(cache=true) when cache does not exist",
			setUp: func(mockCache *httpclientcache.MockClient) {
				mockCache.EXPECT().Get(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(false, nil)
				gock.New("https://www.example.com").
					Get("/").
					Reply(http.StatusOK).
					AddHeader("Content-type", "text/html").
					BodyString("Hello, world!")
				mockCache.EXPECT().Set(
					gomock.Any(),
					gomock.Any(),
					"text/html",
					gomock.Any(),
					gomock.Any(),
				)
			},
			inputUseCache:           true,
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/ query:map[]]]"`,
			},
		},
		{
			desc: "Success(cache=true) when cache exists",
			setUp: func(mockCache *httpclientcache.MockClient) {
				mockCache.EXPECT().Get(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(true, nil)
			},
			inputUseCache:           true,
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines:        []string{},
		},
		{
			desc: "Success(cache=true) when cache storage is down",
			setUp: func(mockCache *httpclientcache.MockClient) {
				mockCache.EXPECT().Get(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(false, errors.New("dummy"))
				gock.New("https://www.example.com").
					Get("/").
					Reply(http.StatusOK).
					AddHeader("Content-type", "text/html").
					BodyString("Hello, world!")
				mockCache.EXPECT().Set(
					gomock.Any(),
					gomock.Any(),
					"text/html",
					gomock.Any(),
					gomock.Any(),
				).Return(errors.New("dummy"))
			},
			inputUseCache:           true,
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=WARN msg="Failed to get cache" err=dummy`,
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/ query:map[]]]"`,
				`level=WARN msg="Failed to set cache" err=dummy`,
			},
		},
		{
			desc: "Error - Unknown http client error",
			setUp: func(mockCache *httpclientcache.MockClient) {
				gock.New("https://www.example.com").
					Get("/").ReplyError(errors.New("dummy"))
			},
			inputUseCache:           false,
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/ query:map[]]]"`,
			},
			expectedError: `Get "https://www.example.com/": dummy`,
		},
		{
			desc: "Error - Http error",
			setUp: func(mockCache *httpclientcache.MockClient) {
				gock.New("https://www.example.com").
					Get("/").
					Reply(http.StatusBadRequest).
					BodyString("Hello, world!")
			},
			inputUseCache:           false,
			inputStatusCodesSuccess: []int{http.StatusOK},
			expectedLogLines: []string{
				`level=INFO msg="" fetcher="map[request:map[host:www.example.com path:/ query:map[]]]"`,
			},
			expectedError: `HTTP error : status=400`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer gock.Off()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cache := httpclientcache.NewMockClient(ctrl)
			tC.setUp(cache)
			cli := HTTPClientWrapperImpl{
				Cli:         http.DefaultClient,
				UseCache:    tC.inputUseCache,
				Cache:       cache,
				CacheOption: &httpclientcache.ClientOption{},
			}
			logger, logBuffer := clog.NewBytesBufferLogger()
			w := bytes.NewBuffer([]byte{})
			err := cli.Do(context.Background(), logger, inputReq, w, tC.inputStatusCodesSuccess)
			test_helper.AssertError(t, tC.expectedError, err)
			clog.AssertLogString(t, tC.expectedLogLines, logBuffer)
		})
	}
}
