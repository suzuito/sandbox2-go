package factorynewfuncs

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/web"
)

var NewFuncsFetcher = []factory.NewFuncFetcher{
	web.NewFetcherHTTP,
	web.NewFetcherHTTPStatic,
	web.NewFetcherHTTPConnpass,
}
