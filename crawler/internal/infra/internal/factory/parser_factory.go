package factory

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type NewFuncParser func(def *crawler.ParserDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Parser, error)

type ParserFactory struct {
	CrawlerFactorySetting *factorysetting.CrawlerFactorySetting
	NewFuncs              []NewFuncParser
}

func (t *ParserFactory) Get(ctx context.Context, def *crawler.ParserDefinition) (crawler.Parser, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, t.CrawlerFactorySetting)
		if err != nil {
			if errors.Is(err, ErrNoMatchedParserID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Parser '%s' is not found in available list", def.ID)
}
