package httprequestcache

import (
	"context"
	"io"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/httprequestcacheimpl"
)

type HTTPRequestCache interface {
	Do(ctx context.Context, req *http.Request) (hit bool, r io.Reader, err error)
}

func NewHTTPRequestCache(
	bucketName string,
	firestoreClient *firestore.Client,
) HTTPRequestCache {
	return &httprequestcacheimpl.HTTPRequestCacheImpl{
		BucketName:      bucketName,
		FirestoreClient: firestoreClient,
	}
}
