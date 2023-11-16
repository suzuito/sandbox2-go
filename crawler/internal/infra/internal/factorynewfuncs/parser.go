package factorynewfuncs

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/parserimpl/goblog"
)

var NewFuncsParser = []factory.NewFuncParser{
	goblog.New,
}
