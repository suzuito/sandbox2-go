package gcp

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/internal/common/test_helper"
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

func TestGetTimeSeriesDataFromFullPathFirestore(t *testing.T) {
	baseCollection := "TestGetTimeSeriesDataFromFullPathFirestore"
	testCases := []struct {
		testhelper.TestCaseForFirestore
		desc          string
		inputFullPath string
		expectedError string
	}{
		{
			desc:          "",
			inputFullPath: "TestGetTimeSeriesDataFromFullPathFirestore/foo",
			TestCaseForFirestore: testhelper.TestCaseForFirestore{
				SetUp: func(ctx context.Context, fcli *firestore.Client) error {
					return testhelper.SetDocuments(
						ctx,
						fcli,
						testhelper.SetDocumentsRef{
							Ref: fcli.Collection(baseCollection).Doc("foo"),
							Data: DummyTimeSeriesData{
								ID:          "foo",
								PublishedAt: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
							},
						},
					)
				},
				TearDown: func(ctx context.Context, fcli *firestore.Client) error {
					return testhelper.DeleteDocuments(ctx, fcli, fcli.Collection(baseCollection))
				},
			},
		},
	}
	for _, tC := range testCases {
		ctx := context.Background()
		tC.Run(ctx, tC.desc, t, func(t *testing.T, fcli *firestore.Client) {
			repo := NewRepository(fcli)
			data := DummyTimeSeriesData{}
			err := repo.GetTimeSeriesDataFromFullPathFirestore(ctx, tC.inputFullPath, &data)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
