package crawler

type CrawlerDefinition struct {
	ID                  CrawlerID
	FetcherDefinition   FetcherDefinition
	ParserDefinition    ParserDefinition
	PublisherDefinition PublisherDefinition
}
