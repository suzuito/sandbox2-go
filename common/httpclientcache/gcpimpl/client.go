package gcpimpl

import (
	"time"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/common/httpclientcache"
	internal_gcpimpl "github.com/suzuito/sandbox2-go/common/httpclientcache/gcpimpl/internal/gcpimpl"
)

func New(
	cli *storage.Client,
	bucket string,
	basePath string,
) httpclientcache.Client {
	return &internal_gcpimpl.Client{
		Cli:      cli,
		Bucket:   bucket,
		NowFunc:  time.Now,
		BasePath: basePath,
	}
}
