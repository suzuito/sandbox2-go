package factorysetting

import "github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper"

type FetcherFactorySetting struct {
	HTTPClientWrapper httpclientwrapper.HTTPClientWrapper
}
