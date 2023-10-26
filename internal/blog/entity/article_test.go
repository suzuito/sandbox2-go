package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/internal/common/test_helper"
)

func TestArticleValidate(t *testing.T) {
	var verr *ValidationError
	testCases := []struct {
		desc          string
		input         Article
		expectedError string
	}{
		{
			desc: "Success",
			input: Article{
				ID:          "001",
				Version:     10,
				Title:       "t001",
				Description: "d001",
				Tags:        []Tag{},
				ArticleSource: ArticleSource{
					ID:      "s001",
					Version: "aaa",
				},
			},
		},
		{
			desc:          "Error",
			input:         Article{},
			expectedError: ".*Error:Field validation.*",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.input.Validate()
			test_helper.AssertError(t, tC.expectedError, err)
			test_helper.AssertErrorAs(t, &verr, err)
			if tC.expectedError == "" {
				assert.Nil(t, err)
			}
		})
	}
}

func TestArticleDateString(t *testing.T) {
	input := Article{
		Date: time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
	}
	assert.Equal(t, "2023-01-02", input.DateString())
}
