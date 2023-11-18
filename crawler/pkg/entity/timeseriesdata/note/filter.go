package note

import (
	"slices"
	"strings"
)

func FuncHasTag(tags []string) func(*NoteArticle) bool {
	return func(article *NoteArticle) bool {
		hasTag := slices.ContainsFunc(article.Tags, func(tag NoteArticleTag) bool {
			name := strings.ToLower(tag.Name)
			for _, tag := range tags {
				if name == tag {
					return true
				}
			}
			return false
		})
		return hasTag
	}
}
