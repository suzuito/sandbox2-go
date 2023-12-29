package factorysetting

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/common/httpclientcache"
)

type FetcherFactorySetting struct {
	HTTPClient            *http.Client
	HTTPClientCacheClient httpclientcache.Client
}
