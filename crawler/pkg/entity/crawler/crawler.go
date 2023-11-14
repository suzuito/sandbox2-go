package crawler

type CrawlerID string

type CrawlerInputData map[string]interface{}

type Crawler struct {
	ID        CrawlerID
	Fetcher   Fetcher
	Parser    Parser
	Publisher Publisher
}
