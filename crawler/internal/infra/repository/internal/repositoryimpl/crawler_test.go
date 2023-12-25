package repositoryimpl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func TestCrawlerRepositoryGetCrawlerDefinition(t *testing.T) {
	testCases := []struct {
		desc          string
		inputCrawlers map[crawler.CrawlerID]*crawler.CrawlerDefinition
		inputID       crawler.CrawlerID
		expected      crawler.CrawlerDefinition
		expectedError string
	}{
		{
			desc: "Success",
			inputCrawlers: map[crawler.CrawlerID]*crawler.CrawlerDefinition{
				"c001": {ID: "c001"},
			},
			inputID:  "c001",
			expected: crawler.CrawlerDefinition{ID: "c001"},
		},
		{
			desc:          "Failed",
			inputCrawlers: map[crawler.CrawlerID]*crawler.CrawlerDefinition{},
			inputID:       "c001",
			expectedError: `CrawlerDefinition\[c001\] is not found`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := CrawlerRepository{
				Crawlers: tC.inputCrawlers,
			}
			def, err := r.GetCrawlerDefinition(context.Background(), tC.inputID)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expected.ID, def.ID)
			}
		})
	}
}

func TestCrawlerRepositoryGetCrawlerStarterSettings(t *testing.T) {
	testCases := []struct {
		desc                         string
		inputCrawlerSettings         []*crawler.CrawlerStarterSetting
		inputCrawlerStarterSettingID crawler.CrawlerStarterSettingID
		expectedLen                  int
		expectedError                string
	}{
		{
			desc: "Success",
			inputCrawlerSettings: []*crawler.CrawlerStarterSetting{
				{ID: "starter001"},
				{ID: "starter002"},
				{ID: "starter001"},
			},
			inputCrawlerStarterSettingID: "starter001",
			expectedLen:                  2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := CrawlerRepository{
				CrawlerSettings: tC.inputCrawlerSettings,
			}
			settings, err := r.GetCrawlerStarterSettings(context.Background(), tC.inputCrawlerStarterSettingID)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedLen, len(settings))
			}
		})
	}
}
