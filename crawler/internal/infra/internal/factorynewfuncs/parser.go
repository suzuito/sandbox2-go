package factorynewfuncs

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/parserimpl/connpass"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/parserimpl/goblog"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/parserimpl/notecontent"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/parserimpl/rss"
)

var NewFuncsParser = []factory.NewFuncParser{
	goblog.New,
	rss.New,
	connpass.New,
	notecontent.New,
}
