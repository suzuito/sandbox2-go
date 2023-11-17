package note

import (
	"slices"
	"strings"
)

func HasGolangTag(article *NoteArticle) bool {
	hasGolangTag := slices.ContainsFunc(article.Tags, func(tag NoteArticleTag) bool {
		name := strings.ToLower(tag.Name)
		if name == "go" || name == "golang" || name == "go言語" {
			return true
		}
		return false
	})
	return hasGolangTag
}
