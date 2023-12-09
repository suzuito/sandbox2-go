package crawler

type CrawlFunctionID string

type DispatchCrawlSetting struct {
	CrawlFunctionIDMapping map[CrawlerID]CrawlFunctionID
	DefaultCrawlFunctionID CrawlFunctionID
}
