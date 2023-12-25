package fetchercacheimpl

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/fetchercacheimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/fetchercache"
)

func NewFetcherCache() fetchercache.Client {
	return &fetchercacheimpl.Client{}
}
