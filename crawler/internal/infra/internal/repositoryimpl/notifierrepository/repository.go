package notifierrepository

import (
	"context"
	"regexp"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

type Repository struct {
	Defs []notifier.NotifierDefinition
}

func (t *Repository) GetNotiferDefinitionsFromDocPathFirestore(
	ctx context.Context,
	fullPath string,
) ([]notifier.NotifierDefinition, error) {
	matchedDefs := []notifier.NotifierDefinition{}
	for _, n := range t.Defs {
		for _, matcher := range n.DocPathFirestoreMatchers {
			matched, err := regexp.MatchString(matcher, fullPath)
			if err != nil {
				return nil, terrors.Wrap(err)
			}
			if !matched {
				continue
			}
			matchedDefs = append(matchedDefs, n)
		}
	}
	return matchedDefs, nil
}
