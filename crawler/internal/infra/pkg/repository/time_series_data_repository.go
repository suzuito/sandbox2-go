package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/repositoryimpl/timeseriesdatarepository"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
)

func NewTimeSeriesDataRepository(
	cli *firestore.Client,
	baseCollection string,
) repository.TimeSeriesDataRepository {
	return &timeseriesdatarepository.Repository{
		Cli:            cli,
		BaseCollection: baseCollection,
	}
}
