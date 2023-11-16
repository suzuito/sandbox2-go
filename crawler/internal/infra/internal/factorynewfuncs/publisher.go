package factorynewfuncs

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/publisherimpl/enqueuecrawl"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/publisherimpl/timeseriesdatarepository"
)

var NewFuncsPublisher = []factory.NewFuncPublisher{
	timeseriesdatarepository.New,
	enqueuecrawl.New,
}
