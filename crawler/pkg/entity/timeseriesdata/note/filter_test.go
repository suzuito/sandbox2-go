package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasGolangTag(t *testing.T) {
	testCases := []struct {
		desc         string
		inputArticle NoteArticle
		expected     bool
	}{
		{
			desc: "go",
			inputArticle: NoteArticle{
				Tags: []NoteArticleTag{
					{Name: "GO"},
				},
			},
			expected: true,
		},
		{
			desc: "golang",
			inputArticle: NoteArticle{
				Tags: []NoteArticleTag{
					{Name: "Golang"},
				},
			},
			expected: true,
		},
		{
			desc: "golang",
			inputArticle: NoteArticle{
				Tags: []NoteArticleTag{
					{Name: "Go言語"},
				},
			},
			expected: true,
		},
		{
			desc: "GOO",
			inputArticle: NoteArticle{
				Tags: []NoteArticleTag{
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
