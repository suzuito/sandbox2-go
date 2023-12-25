package repositoryimpl

import (
	"context"
	"errors"
	"net/http"

	"cloud.google.com/go/firestore"
)

type HTTPClientCacheRepositoryImpl struct {
	Cli            *firestore.Client
	BaseCollection string
}

func (t *HTTPClientCacheRepositoryImpl) Get(ctx context.Context, req *http.Request) (bool, []byte, error) {
	return false, nil, errors.New("not implemented")
}

func (t *HTTPClientCacheRepositoryImpl) Set(ctx context.Context, req *http.Request, resBody []byte) error {
	return errors.New("not implemented")
}
