package httprequestcacheimpl

import (
	"context"
	"io"
	"net/http"

	"cloud.google.com/go/firestore"
)

type HTTPRequestCacheImpl struct {
	BucketName      string
	FirestoreClient *firestore.Client
}

func (t *HTTPRequestCacheImpl) Do(ctx context.Context, req *http.Request) (bool, io.Reader, error) {
	// TODO implement
	return false, nil, nil
}
