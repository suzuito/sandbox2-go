package repository

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/repository/internal/repositoryimpl"
)

type HTTPClientCacheRepository interface {
	Get(ctx context.Context, req *http.Request) (bool, []byte, error)
	Set(ctx context.Context, req *http.Request, resBody []byte) error
}

func NewHTTPClientCacheRepository(
	cli *firestore.Client,
	baseCollection string,
) *repositoryimpl.HTTPClientCacheRepositoryImpl {
	return &repositoryimpl.HTTPClientCacheRepositoryImpl{
		Cli:            cli,
		BaseCollection: baseCollection,
	}
}
