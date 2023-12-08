package crawler

type CrawlerStarterSettingID string

type CrawlerStarterSetting struct {
	ID               CrawlerStarterSettingID
	CrawlerID        CrawlerID
	CrawlerInputData CrawlerInputData
}
