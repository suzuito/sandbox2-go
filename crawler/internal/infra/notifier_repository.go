package infra

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/repositoryimpl/notifierrepository"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

func NewNotifierRepository(defs []notifier.NotifierDefinition) repository.NotifierRepository {
	return &notifierrepository.Repository{
		Defs: defs,
	}
}
