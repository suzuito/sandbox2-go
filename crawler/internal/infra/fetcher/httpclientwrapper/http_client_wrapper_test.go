package httpclientwrapper

import (
	"testing"

	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestNewHTTPClientWrapperFromArgument(t *testing.T) {
	NewHTTPClientWrapperFromArgument(
		&crawler.FetcherDefinition{},
		&factorysetting.CrawlerFactorySetting{},
	)
}
