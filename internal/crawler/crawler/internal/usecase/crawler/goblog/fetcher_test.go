package goblog

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/internal/common/test_helper"
)

func TestFetch(t *testing.T) {
	testCases := []struct {
		desc          string
		setUp         func()
		expectedW     string
		expectedError string
	}{
		{
			desc:      "200",
			expectedW: "hoge",
			setUp: func() {
				gock.New("https://go.dev").
					Get("/blog").
					Reply(200).
					BodyString("hoge")
			},
		},
		{
			desc: "Conn error",
			setUp: func() {
				gock.New("https://go.dev").
					Get("/blog").
					ReplyError(fmt.Errorf("dummy error"))
			},
			expectedError: "dummy error",
		},
		{
			desc: "HTTP error",
			setUp: func() {
				gock.New("https://go.dev").
					Get("/blog").
					Reply(404).
					BodyString("hoge")
			},
			expectedError: "HTTP error is occured code=404",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer gock.Off()
			ctx := context.Background()
			tC.setUp()
			fetcher := Fetcher{cliHTTP: http.DefaultClient}
			w := bytes.NewBufferString("")
			err := fetcher.Fetch(ctx, w)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedW, w.String())
			}
		})
	}
}
