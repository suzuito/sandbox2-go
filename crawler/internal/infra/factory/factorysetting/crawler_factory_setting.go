package factorysetting

type CrawlerFactorySetting struct {
	FetcherFactorySetting   FetcherFactorySetting
	ParserFactorySetting    ParserFactorySetting
	PublisherFactorySetting PublisherFactorySetting
}
