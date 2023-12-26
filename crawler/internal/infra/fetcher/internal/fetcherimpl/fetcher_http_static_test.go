package fetcherimpl

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
)

func TestFetcherHTTPStaticDo(t *testing.T) {
	testCases := []struct {
		desc                    string
		mockCli                 MockHTTPClientWrapper
		inputReqFunc            func() *http.Request
		inputStatusCodesSuccess []int
		expectedError           string
	}{
		{
			desc: "Success",
			mockCli: MockHTTPClientWrapper{
				ExpectedMethod: http.MethodGet,
				ExpectedURL:    mustURL("https://www.example.com/hoge/"),
			},
			inputReqFunc: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "https://www.example.com/hoge/", nil)
				return req
			},
			inputStatusCodesSuccess: []int{http.StatusOK},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := FetcherHTTPStatic{
				Cli:                &tC.mockCli,
				Req:                tC.inputReqFunc(),
				StatusCodesSuccess: tC.inputStatusCodesSuccess,
			}
			w := bytes.NewBuffer([]byte{})
			logger, _ := newMockLogger()
			err := f.Do(context.Background(), logger, w, nil)
			test_helper.AssertError(t, tC.expectedError, err)
			tC.mockCli.Assert(t)
		})
	}
}
