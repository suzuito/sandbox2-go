package httpclientcache

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyGenKeyString(t *testing.T) {
	inputRequest1 := http.Request{
		URL: &url.URL{
			Scheme:   "https",
			Host:     "www.example.com",
			Path:     "/hoge/fuga",
			RawQuery: "key1=a&key1=b&key1=c&key2=d&key2=e&key2=f",
		},
	}
	inputRequest2 := http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "www.example.com",
			Path:   "/hoge/fuga",
		},
	}
	inputRequest3 := http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "www.example.com",
			Path:   "hoge/fuga",
		},
	}
	testCases := []struct {
		desc           string
		input          KeyGen
		inputRequest   http.Request
		expectedString string
	}{
		{
			input:          KeyGen{},
			inputRequest:   inputRequest1,
			expectedString: "",
		},
		{
			input: KeyGen{
				IncludeProtocol: true,
			},
			inputRequest:   inputRequest1,
			expectedString: "https://",
		},
		{
			input: KeyGen{
				IncludeHost: true,
			},
			inputRequest:   inputRequest1,
			expectedString: "www.example.com",
		},
		{
			input: KeyGen{
				IncludePath: true,
			},
			inputRequest:   inputRequest1,
			expectedString: "/hoge/fuga",
		},
		{
			input: KeyGen{
				IncludePath: true,
			},
			inputRequest:   inputRequest3,
			expectedString: "/hoge/fuga",
		},
		{
			input: KeyGen{
				IncludeQuery: true,
				QueryKeys:    []string{"key1", "key2"},
			},
			inputRequest:   inputRequest1,
			expectedString: "?key1=a&key1=b&key1=c&key2=d&key2=e&key2=f",
		},
		{
			input: KeyGen{
				IncludeQuery: true,
				QueryKeys:    []string{"key2"},
			},
			inputRequest:   inputRequest1,
			expectedString: "?key2=d&key2=e&key2=f",
		},
		{
			input: KeyGen{
				IncludeProtocol: true,
				IncludeHost:     true,
				IncludePath:     true,
				IncludeQuery:    true,
				QueryKeys:       []string{"key1", "key2"},
			},
			inputRequest:   inputRequest1,
			expectedString: "https://www.example.com/hoge/fuga?key1=a&key1=b&key1=c&key2=d&key2=e&key2=f",
		},
		{
			input: KeyGen{
				IncludeProtocol: true,
				IncludeHost:     true,
				IncludePath:     true,
				IncludeQuery:    true,
				QueryKeys:       []string{"key2"},
			},
			inputRequest:   inputRequest1,
			expectedString: "https://www.example.com/hoge/fuga?key2=d&key2=e&key2=f",
		},
		{
			input: KeyGen{
				IncludeProtocol: true,
				IncludeHost:     true,
				IncludePath:     true,
				IncludeQuery:    true,
				QueryKeys:       []string{"foo", "bar"},
			},
			inputRequest:   inputRequest1,
			expectedString: "https://www.example.com/hoge/fuga",
		},
		{
			input: KeyGen{
				IncludeProtocol: true,
				IncludeHost:     true,
				IncludePath:     true,
				IncludeQuery:    true,
				QueryKeys:       []string{"key2"},
			},
			inputRequest:   inputRequest2,
			expectedString: "https://www.example.com/hoge/fuga",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expectedString, tC.input.KeyString(&tC.inputRequest))
		})
	}
}
