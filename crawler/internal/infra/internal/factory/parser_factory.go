package factory

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type NewFuncParserArgument struct {
}
type NewFuncParser func(def *crawler.ParserDefinition, arg *NewFuncParserArgument) (crawler.Parser, error)

type ParserFactory struct {
	NewFuncs []NewFuncParser
}

func (t *ParserFactory) Get(ctx context.Context, def *crawler.ParserDefinition) (crawler.Parser, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, &NewFuncParserArgument{})
		if err != nil {
			if errors.Is(err, ErrNoMatchedFetcherID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Parser '%s' is not found in available list", def.ID)
}
