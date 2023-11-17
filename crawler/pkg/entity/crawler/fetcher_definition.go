package crawler

import "github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"

type FetcherDefinition struct {
	ID       FetcherID
	Argument argument.ArgumentDefinition
}
