package cweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	testCases := []struct {
		desc            string
		inputQueryKey   string
		inputQueryValue string
		inputDefault    int
		expected        int
	}{
		{
			desc:            "Convertable query",
			inputQueryKey:   "q",
			inputQueryValue: "1",
			inputDefault:    100,
			expected:        1,
		},
		{
			desc:         "Empty query",
			inputDefault: 100,
			expected:     100,
		},
		{
			desc:            "Unconvertable (str)",
			inputQueryKey:   "q",
			inputQueryValue: "hoge",
			inputDefault:    100,
			expected:        100,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "", nil)
			if tC.inputQueryKey != "" {
				req.URL.RawQuery = fmt.Sprintf("%s=%s", tC.inputQueryKey, tC.inputQueryValue)
			}
			ctx := gin.Context{
				Request: req,
			}
			real := DefaultQueryAsInt(&ctx, tC.inputQueryKey, tC.inputDefault)
			assert.Equal(t, tC.expected, real)
		})
	}
}

func TestInt64(t *testing.T) {
	testCases := []struct {
		desc            string
		inputQueryKey   string
		inputQueryValue string
		inputDefault    int64
		expected        int64
	}{
		{
			desc:            "Convertable query",
			inputQueryKey:   "q",
			inputQueryValue: "1",
			inputDefault:    100,
			expected:        1,
		},
		{
			desc:         "Empty query",
			inputDefault: 100,
			expected:     100,
		},
		{
			desc:            "Unconvertable (str)",
			inputQueryKey:   "q",
			inputQueryValue: "hoge",
			inputDefault:    100,
			expected:        100,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "", nil)
			if tC.inputQueryKey != "" {
				req.URL.RawQuery = fmt.Sprintf("%s=%s", tC.inputQueryKey, tC.inputQueryValue)
			}
			ctx := gin.Context{
				Request: req,
			}
			real := DefaultQueryAsInt64(&ctx, tC.inputQueryKey, tC.inputDefault)
			assert.Equal(t, tC.expected, real)
		})
	}
}

func TestFloat32(t *testing.T) {
	testCases := []struct {
		desc            string
		inputQueryKey   string
		inputQueryValue string
		inputDefault    float32
		expected        float32
	}{
		{
			desc:            "Convertable query",
			inputQueryKey:   "q",
			inputQueryValue: "1",
			inputDefault:    100,
			expected:        1,
		},
		{
			desc:         "Empty query",
			inputDefault: 100,
			expected:     100,
		},
		{
			desc:            "Unconvertable (str)",
			inputQueryKey:   "q",
			inputQueryValue: "hoge",
			inputDefault:    100,
			expected:        100,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "", nil)
			if tC.inputQueryKey != "" {
				req.URL.RawQuery = fmt.Sprintf("%s=%s", tC.inputQueryKey, tC.inputQueryValue)
			}
			ctx := gin.Context{
				Request: req,
			}
			real := DefaultQueryAsFloat32(&ctx, tC.inputQueryKey, tC.inputDefault)
			assert.Equal(t, tC.expected, real)
		})
	}
}

func TestFloat64(t *testing.T) {
	testCases := []struct {
		desc            string
		inputQueryKey   string
		inputQueryValue string
		inputDefault    float64
		expected        float64
	}{
		{
			desc:            "Convertable query",
			inputQueryKey:   "q",
			inputQueryValue: "1",
			inputDefault:    100,
			expected:        1,
		},
		{
			desc:         "Empty query",
			inputDefault: 100,
			expected:     100,
		},
		{
			desc:            "Unconvertable (str)",
			inputQueryKey:   "q",
			inputQueryValue: "hoge",
			inputDefault:    100,
			expected:        100,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "", nil)
			if tC.inputQueryKey != "" {
				req.URL.RawQuery = fmt.Sprintf("%s=%s", tC.inputQueryKey, tC.inputQueryValue)
			}
			ctx := gin.Context{
				Request: req,
			}
			real := DefaultQueryAsFloat64(&ctx, tC.inputQueryKey, tC.inputDefault)
			assert.Equal(t, tC.expected, real)
		})
	}
}

func TestBool(t *testing.T) {
	testCases := []struct {
		desc            string
		inputQueryKey   string
		inputQueryValue string
		inputDefault    bool
		expected        bool
	}{
		{
			desc:            "Convertable query",
			inputQueryKey:   "q",
			inputQueryValue: "true",
			inputDefault:    false,
			expected:        true,
		},
		{
			desc:         "Empty query",
			inputDefault: false,
			expected:     false,
		},
		{
			desc:            "Unconvertable (str)",
			inputQueryKey:   "q",
			inputQueryValue: "hoge",
			inputDefault:    false,
			expected:        false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "", nil)
			if tC.inputQueryKey != "" {
				req.URL.RawQuery = fmt.Sprintf("%s=%s", tC.inputQueryKey, tC.inputQueryValue)
			}
			ctx := gin.Context{
				Request: req,
			}
			real := DefaultQueryAsBool(&ctx, tC.inputQueryKey, tC.inputDefault)
			assert.Equal(t, tC.expected, real)
		})
	}
}
