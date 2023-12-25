package parserimpl

import (
	"context"
	"io"
	"strings"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

type NoteContentParser struct {
	FilterByTag []string
}

func (t *NoteContentParser) ID() crawler.ParserID {
	return "notecontent"
}

func (t *NoteContentParser) Do(ctx context.Context, r io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	p := note.Parser{}
	d, err := p.Parse(ctx, r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if !note.FuncHasTag(t.FilterByTag)(d) {
		return []timeseriesdata.TimeSeriesData{}, nil
	}
	blogFeed := timeseriesdata.TimeSeriesDataBlogFeed{}
	blogFeed.ID = timeseriesdata.TimeSeriesDataID(strings.ReplaceAll(strings.ReplaceAll(d.URL, ":", "-"), "/", "-"))
	blogFeed.Title = d.Title
	blogFeed.Summary = d.Description
	blogFeed.PublishedAt = d.PublishedAt
	blogFeed.ArticleContent = d.ArticleContent
	blogFeed.Thumbnail = &timeseriesdata.TimeSeriesDataBlogFeedThumbnail{
		ImageURL: d.ImageURL,
	}
	blogFeed.URL = d.URL
	blogFeed.Author = &timeseriesdata.TimeSeriesDataBlogFeedAuthor{
		URL:  d.AuthorURL,
		Name: d.AuthorName,
	}
	return []timeseriesdata.TimeSeriesData{
		&blogFeed,
	}, nil
}
