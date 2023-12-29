package gcpimpl

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/httpclientcache"
	"github.com/suzuito/sandbox2-go/common/test_helper"
)

// GCPに繋いだテスト書いてみる
func TestSetAndGet(t *testing.T) {
	var testbucket = "suzuito-minilla-ut"
	var testBasePath = "common.httpclientcache.gcpimpl.internal.gcpimpl.TestSetAndGet"
	ctx := context.Background()
	scli, err := storage.NewClient(ctx)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	testCases := []struct {
		desc string
		do   func(t *testing.T)
	}{
		{
			desc: "Hit (KeyGen all nil)",
			do: func(t *testing.T) {
				cli := Client{
					Bucket:   testbucket,
					BasePath: testBasePath,
					Cli:      scli,
					NowFunc:  time.Now,
				}
				inputBody := bytes.NewBufferString("hoge")
				inputRequest := http.Request{
					URL: &url.URL{
						Scheme:   "https",
						Host:     "www.example.com",
						Path:     "/hoge/fuga",
						RawQuery: "key1=a&key1=b&key1=c&key2=d&key2=e&key2=f",
					},
				}
				inputOption := httpclientcache.ClientOption{
					KeyGen:    &httpclientcache.KeyGen{},
					TTLInDays: 100,
				}
				err = cli.Set(ctx, &inputRequest, "text/html", inputBody, &inputOption)
				if err != nil {
					t.Errorf("%+v", err)
					return
				}
				body := bytes.NewBuffer([]byte{})
				hit, err := cli.Get(ctx, &inputRequest, body, &inputOption)
				if err != nil {
					t.Errorf("%+v", err)
					return
				}
				assert.True(t, hit)
				assert.Equal(t, "hoge", body.String())
			},
		},
		{
			desc: "Hit",
			do: func(t *testing.T) {
				cli := Client{
					Bucket:   testbucket,
					BasePath: testBasePath,
					Cli:      scli,
					NowFunc:  time.Now,
				}
				inputBody := bytes.NewBufferString("hoge")
				inputRequest := http.Request{
					URL: &url.URL{
						Scheme:   "https",
						Host:     "www.example.com",
						Path:     "/hoge/fuga",
						RawQuery: "key1=a&key1=b&key1=c&key2=d&key2=e&key2=f",
					},
				}
				inputOption := httpclientcache.ClientOption{
					KeyGen: &httpclientcache.KeyGen{
						IncludeProtocol: true,
						IncludeHost:     true,
						IncludePath:     true,
						IncludeQuery:    true,
						QueryKeys:       []string{"key2"},
					},
					TTLInDays: 100,
				}
				err = cli.Set(ctx, &inputRequest, "text/html", inputBody, &inputOption)
				if err != nil {
					t.Errorf("%+v", err)
					return
				}
				body := bytes.NewBuffer([]byte{})
				hit, err := cli.Get(ctx, &inputRequest, body, &inputOption)
				if err != nil {
					t.Errorf("%+v", err)
					return
				}
				assert.True(t, hit)
				assert.Equal(t, "hoge", body.String())
			},
		},
		{
			desc: "Not hit (Object does not exist)",
			do: func(t *testing.T) {
				cli := Client{
					Bucket:   testbucket,
					BasePath: testBasePath,
					Cli:      scli,
					NowFunc:  time.Now,
				}
				inputRequest := http.Request{
					URL: &url.URL{
						Scheme:   "https",
						Host:     "www.example.com",
						Path:     "/hoge/fuga",
						RawQuery: "key1=a&key1=b&key1=c&key2=d&key2=e&key2=f",
					},
				}
				inputOption := httpclientcache.ClientOption{
					KeyGen: &httpclientcache.KeyGen{
						IncludeProtocol: true,
						IncludeHost:     true,
						IncludePath:     true,
						IncludeQuery:    true,
						QueryKeys:       []string{"key2"},
					},
					TTLInDays: 100,
				}
				body := bytes.NewBuffer([]byte{})
				hit, err := cli.Get(ctx, &inputRequest, body, &inputOption)
				if err != nil {
					t.Errorf("%+v", err)
					return
				}
				assert.False(t, hit)
			},
		},
		{
			desc: "Not hit (Expired)",
			do: func(t *testing.T) {
				cli := Client{
					Bucket:   testbucket,
					BasePath: testBasePath,
					Cli:      scli,
					NowFunc:  time.Now,
				}
				inputRequest := http.Request{
					URL: &url.URL{
						Scheme:   "https",
						Host:     "www.example.com",
						Path:     "/hoge/fuga",
						RawQuery: "key1=a&key1=b&key1=c&key2=d&key2=e&key2=f",
					},
				}
				inputBody := bytes.NewBufferString("hoge")
				inputOption := httpclientcache.ClientOption{
					KeyGen: &httpclientcache.KeyGen{
						IncludeProtocol: true,
						IncludeHost:     true,
						IncludePath:     true,
						IncludeQuery:    true,
						QueryKeys:       []string{"key2"},
					},
					TTLInDays: 0,
				}
				err = cli.Set(ctx, &inputRequest, "text/html", inputBody, &inputOption)
				if err != nil {
					t.Errorf("%+v", err)
					return
				}
				body := bytes.NewBuffer([]byte{})
				hit, err := cli.Get(ctx, &inputRequest, body, &inputOption)
				if err != nil {
					t.Errorf("%+v", err)
					return
				}
				assert.False(t, hit)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			defer func() {
				if err := test_helper.DeleteGoogleCloudStorageObjects(ctx, scli, testbucket, testBasePath); err != nil {
					t.Errorf("%+v", err)
				}
			}()
			tc.do(t)
		})
	}

}
