package factorysetting

import "net/http"

type FetcherFactorySetting struct {
	HTTPClient *http.Client
}
