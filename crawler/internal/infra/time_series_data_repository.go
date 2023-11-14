package infra

import (
	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/gcp"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
)

func NewTimeSeriesDataRepository(
	cli *firestore.Client,
	baseCollection string,
) repository.TimeSeriesDataRepository {
	return &gcp.TimeSeriesDataRepository{
		Cli:            cli,
		BaseCollection: baseCollection,
	}
}
