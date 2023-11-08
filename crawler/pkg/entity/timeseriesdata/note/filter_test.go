package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasGolangTag(t *testing.T) {
	testCases := []struct {
		desc         string
		inputArticle TimeSeriesDataNoteArticle
		expected     bool
	}{
		{
			desc: "go",
			inputArticle: TimeSeriesDataNoteArticle{
				Tags: []TimeSeriesDataNoteArticleTag{
					{Name: "GO"},
				},
			},
			expected: true,
		},
		{
			desc: "golang",
			inputArticle: TimeSeriesDataNoteArticle{
				Tags: []TimeSeriesDataNoteArticleTag{
					{Name: "Golang"},
				},
			},
			expected: true,
		},
		{
			desc: "golang",
			inputArticle: TimeSeriesDataNoteArticle{
				Tags: []TimeSeriesDataNoteArticleTag{
					{Name: "Go言語"},
				},
			},
			expected: true,
		},
		{
			desc: "GOO",
			inputArticle: TimeSeriesDataNoteArticle{
				Tags: []TimeSeriesDataNoteArticleTag{
					{Name: "GOO"},
				},
			},
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(
				t,
				tC.expected,
				HasGolangTag(&tC.inputArticle),
			)
		})
	}
}
