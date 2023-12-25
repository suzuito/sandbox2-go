package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/repository/internal/repositoryimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
)

func NewTimeSeriesDataRepository(
	cli *firestore.Client,
	baseCollection string,
) repository.TimeSeriesDataRepository {
	return &repositoryimpl.TimeSeriesDataRepository{
		Cli:            cli,
		BaseCollection: baseCollection,
	}
}
