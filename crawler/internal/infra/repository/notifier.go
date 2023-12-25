package repository

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/repository/internal/repositoryimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

func NewNotifierRepository(defs []notifier.NotifierDefinition) repository.NotifierRepository {
	return &repositoryimpl.NotifierRepository{
		Defs: defs,
	}
}
