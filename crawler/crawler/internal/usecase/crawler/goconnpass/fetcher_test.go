package goconnpass

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
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
				gock.New("https://connpass.com").
					Get("/api/v1/event/").
					Reply(200).
					BodyString("hoge")
			},
		},
		{
			desc: "Conn error",
			setUp: func() {
				gock.New("https://connpass.com").
					Get("/api/v1/event/").
					ReplyError(fmt.Errorf("dummy error"))
			},
			expectedError: "dummy error",
		},
		{
			desc: "HTTP error",
			setUp: func() {
				gock.New("https://connpass.com").
					Get("/api/v1/event/").
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
			fetcher := Fetcher{CliHTTP: http.DefaultClient}
			w := bytes.NewBufferString("")
			err := fetcher.Fetch(ctx, w)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedW, w.String())
			}
		})
	}
}
