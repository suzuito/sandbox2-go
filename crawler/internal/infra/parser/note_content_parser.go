package parser

import (
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/parser/internal/parserimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewNoteContentParser(def *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
	parser := parserimpl.NoteContentParser{}
	if def.ID != parser.ID() {
		return nil, factoryerror.ErrNoMatchedParserID
	}
	filterByTags, err := argument.GetFromArgumentDefinition[[]string](def.Argument, "FilterByTags")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	parser.FilterByTag = filterByTags
	return &parser, nil
}
