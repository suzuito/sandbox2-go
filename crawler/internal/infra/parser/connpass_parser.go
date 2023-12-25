package parser

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/parser/internal/parserimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewConnpassParser(def *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
	parser := parserimpl.ConnpassParser{}
	if def.ID != parser.ID() {
		return nil, factoryerror.ErrNoMatchedParserID
	}
	return &parser, nil
}
