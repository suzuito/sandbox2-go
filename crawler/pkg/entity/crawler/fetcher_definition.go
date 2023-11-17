package crawler

import "github.com/suzuito/sandbox2-go/crawler/pkg/argument"

type FetcherDefinition struct {
	ID       FetcherID
	Argument argument.ArgumentDefinition
}
