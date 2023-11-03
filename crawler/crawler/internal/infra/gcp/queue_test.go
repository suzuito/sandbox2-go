package gcp

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/testhelper"
)

func TestPublishCrawlEvent(t *testing.T) {
	topicIDCrawlEvent := "topic-TestPublishCrawlEvent"
	testCases := []struct {
		desc                  string
		inputHasTraceID       bool
		inputCrawlerID        crawler.CrawlerID
		inputCrawlerInputData crawler.CrawlerInputData
		testhelper.TestCaseForPubSub
		expectedError string
	}{
		{
			desc:            "",
			inputHasTraceID: true,
			inputCrawlerID:  "hoge",
			inputCrawlerInputData: crawler.CrawlerInputData{
				"foo": "bar",
			},
			TestCaseForPubSub: testhelper.TestCaseForPubSub{
				SetUp: func(ctx context.Context, fcli *pubsub.Client) error {
					_, err := fcli.CreateTopic(ctx, topicIDCrawlEvent)
					if err != nil {
						return err
					}
					return nil
				},
				TearDown: func(ctx context.Context, fcli *pubsub.Client) error {
					return fcli.Topic(topicIDCrawlEvent).Delete(ctx)
				},
			},
		},
	}
	for _, tC := range testCases {
		ctx := context.Background()
		tC.Run(ctx, tC.desc, t, func(t *testing.T, pcli *pubsub.Client) {
			queue := NewQueue(pcli, topicIDCrawlEvent)
			if tC.inputHasTraceID {
				ctx = context.WithValue(ctx, "traceId", "foo")
			}
			err := queue.PublishCrawlEvent(ctx, tC.inputCrawlerID, tC.inputCrawlerInputData)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}

func TestRecieveCrawlEvent(t *testing.T) {
	ctx := context.Background()
	pcli, err := testhelper.NewPubSubClient(ctx)
	if err != nil {
		t.Errorf("failed to NewPubSubClient : %+v", err)
		t.Fail()
		return
	}
	queue := NewQueue(pcli, "topic-TestPublishCrawlEvent")
	crawlerID, crawlerInputData, err := queue.RecieveCrawlEvent(ctx, []byte(`{"CrawlID":"c1"}`))
	assert.Nil(t, err)
	assert.Equal(t, crawler.CrawlerID("c1"), crawlerID)
	assert.Nil(t, crawlerInputData)
	crawlerID, crawlerInputData, err = queue.RecieveCrawlEvent(ctx, []byte(`{"CrawlID":"c1","CrawlerInputData":{"foo":"bar"}}`))
	assert.Nil(t, err)
	assert.Equal(t, crawler.CrawlerID("c1"), crawlerID)
	assert.Equal(t, crawler.CrawlerInputData{"foo": "bar"}, crawlerInputData)
	_, _, err = queue.RecieveCrawlEvent(ctx, []byte(`aaa`))
	assert.NotNil(t, err)
}
