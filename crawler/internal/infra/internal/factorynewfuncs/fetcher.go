package factorynewfuncs

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/fetcherimpl"
)

var NewFuncsFetcher = []factory.NewFuncFetcher{
	fetcherimpl.NewFetcherHTTP,
	fetcherimpl.NewFetcherHTTPStatic,
	fetcherimpl.NewFetcherHTTPConnpass,
}
