package gcp

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/internal/crawler/testhelper"
)

type DummyTimeSeriesData struct {
	ID          string
	PublishedAt time.Time
}

func (t *DummyTimeSeriesData) GetID() timeseriesdata.TimeSeriesDataID {
	return timeseriesdata.TimeSeriesDataID(t.ID)
}
func (t *DummyTimeSeriesData) GetPublishedAt() time.Time {
	return t.PublishedAt
}

func TestSetTimeSeriesData(t *testing.T) {
	baseCollection := "CrawlerTest"
	ctx := context.Background()
	testCases := []struct {
		testhelper.TestCaseForFirestore
		desc                string
		inputCrawlerID      crawler.CrawlerID
		inputTimeSeriesData []timeseriesdata.TimeSeriesData
		expectedError       string
	}{
		{
			desc: `
			Set 3 data
			`,
			inputCrawlerID: "dummy_crawler_id",
			inputTimeSeriesData: []timeseriesdata.TimeSeriesData{
				&DummyTimeSeriesData{ID: "d1", PublishedAt: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC)},
				&DummyTimeSeriesData{ID: "d2", PublishedAt: time.Date(2000, 1, 1, 1, 1, 2, 0, time.UTC)},
				&DummyTimeSeriesData{ID: "d3", PublishedAt: time.Date(2000, 1, 1, 1, 1, 3, 0, time.UTC)},
			},
			TestCaseForFirestore: testhelper.TestCaseForFirestore{
				Assert: func(ctx context.Context, fcli *firestore.Client, ass *testhelper.FirestoreAssertion) error {
					ass.EqualsDoc(ctx, t, fcli.Doc(fmt.Sprintf("CrawlerTest/TimeSeriesData/dummy_crawler_id/d1")), map[string]interface{}{
						"ID":          "d1",
						"PublishedAt": time.Date(2000, time.January, 1, 1, 1, 1, 0, time.UTC),
					})
					ass.EqualsDoc(ctx, t, fcli.Doc(fmt.Sprintf("CrawlerTest/TimeSeriesData/dummy_crawler_id/d2")), map[string]interface{}{
						"ID":          "d2",
						"PublishedAt": time.Date(2000, time.January, 1, 1, 1, 2, 0, time.UTC),
					})
					ass.EqualsDoc(ctx, t, fcli.Doc(fmt.Sprintf("CrawlerTest/TimeSeriesData/dummy_crawler_id/d3")), map[string]interface{}{
						"ID":          "d3",
						"PublishedAt": time.Date(2000, time.January, 1, 1, 1, 3, 0, time.UTC),
					})
					return nil
				},
				TearDown: func(ctx context.Context, fcli *firestore.Client) error {
					testhelper.DeleteDocuments(ctx, fcli, fcli.Collection(baseCollection))
					return nil
				},
			},
		},
		{
			desc: `
			Set 0 data
			`,
			inputCrawlerID:       "dummy_crawler_id",
			inputTimeSeriesData:  []timeseriesdata.TimeSeriesData{},
			TestCaseForFirestore: testhelper.TestCaseForFirestore{},
		},
	}
	for _, tC := range testCases {
		tC.Run(ctx, tC.desc, t, func(t *testing.T, fcli *firestore.Client) {
			defer testhelper.DeleteDocuments(ctx, fcli, fcli.Collection(baseCollection))
			repo := NewRepository(fcli, baseCollection)
			err := repo.SetTimeSeriesData(
				ctx,
				tC.inputCrawlerID,
				tC.inputTimeSeriesData...,
			)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
