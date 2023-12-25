package factorysetting

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/fetchercache"
)

type FetcherFactorySetting struct {
	HTTPClient         *http.Client
	FetcherCacheClient fetchercache.Client
}
