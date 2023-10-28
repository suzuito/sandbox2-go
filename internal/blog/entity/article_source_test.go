package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
)

func TestArticleSourceMetaScan(t *testing.T) {
	testCases := []struct {
		desc          string
		inputSrc      any
		expectedURL   string
		expectedError string
	}{
		{
			desc:        "Success",
			inputSrc:    []byte(`{"URL":"https://www.example.com"}`),
			expectedURL: "https://www.example.com",
		},
		{
			desc:          `Error/any is no bytes`,
			inputSrc:      1,
			expectedError: "Cannot scan 1",
		},
		{
			desc:          `Error/not json`,
			inputSrc:      []byte(`hoge`),
			expectedError: "^invalid character",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			input := ArticleSourceMeta{}
			err := input.Scan(tC.inputSrc)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedURL, input.URL)
			}
		})
	}
}

func TestArticleSourceMetaValue(t *testing.T) {
	input := ArticleSourceMeta{
		URL: "https://www.example.com",
	}
	v, err := input.Value()
	vv := v.([]byte)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"URL":"https://www.example.com"}`, string(vv))
}
