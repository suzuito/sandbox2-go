package note

import (
	"strings"
	"time"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

// Parser for article on note (note.com)
// ex) https://note.com/knowledgework/n/n46b7881a16a6
type TimeSeriesDataNoteArticle struct {
	URL            string
	ArticleContent string
	PublishedAt    time.Time
	Tags           []TimeSeriesDataNoteArticleTag
}

func (t *TimeSeriesDataNoteArticle) GetID() timeseriesdata.TimeSeriesDataID {
	return timeseriesdata.TimeSeriesDataID(
		strings.ReplaceAll(strings.ReplaceAll(t.URL, ":", "-"), "/", "-"),
	)
}

func (t *TimeSeriesDataNoteArticle) GetPublishedAt() time.Time {
	return t.PublishedAt
}

type TimeSeriesDataNoteArticleTag struct {
	Name string
}
