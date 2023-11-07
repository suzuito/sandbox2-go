package knowledgeworkblog

import (
	"slices"
	"strings"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/notecontent"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

const CrawlerID crawler.CrawlerID = "knowledgeworkblog"

func NewCrawler(
	repository repository.Repository,
	fetcher fetcher.FetcherHTTP,
) crawler.Crawler {
	return notecontent.NewCrawler(
		CrawlerID,
		repository,
		fetcher,
		func(article *note.TimeSeriesDataNoteArticle) bool {
			hasGolangTag := slices.ContainsFunc(article.Tags, func(tag note.TimeSeriesDataNoteArticleTag) bool {
				name := strings.ToLower(tag.Name)
				if name == "go" || name == "golang" || name == "go言語" {
					return true
				}
				return false
			})
			return hasGolangTag
		},
	)
}
