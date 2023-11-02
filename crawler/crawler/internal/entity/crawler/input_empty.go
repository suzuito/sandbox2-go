package crawler

type InputEmpty struct {
	CrawlerID CrawlerID
}

func (t *InputEmpty) GetCrawlerID() CrawlerID {
	return t.CrawlerID
}
