package note

import (
	"time"
)

// Parser for article on note (note.com)
// ex) https://note.com/knowledgework/n/n46b7881a16a6
type NoteArticle struct {
	URL            string
	Title          string
	Description    string
	ArticleContent string
	ImageURL       string
	AuthorURL      string
	AuthorName     string
	PublishedAt    time.Time
	Tags           []NoteArticleTag
}

type NoteArticleTag struct {
	Name string
}
