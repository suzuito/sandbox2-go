package gcp

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type TimeSeriesDataRepository struct {
	Cli            *firestore.Client
	BaseCollection string
}

func (t *TimeSeriesDataRepository) firestoreDocTimeSeriesData(
	timeSeriesDataBaseID timeseriesdata.TimeSeriesDataBaseID,
	id timeseriesdata.TimeSeriesDataID,
) *firestore.DocumentRef {
	return t.Cli.Doc(fmt.Sprintf("%s/TimeSeriesData/%s/%s", t.BaseCollection, timeSeriesDataBaseID, id))
}

func (t *TimeSeriesDataRepository) SetTimeSeriesData(
	ctx context.Context,
	timeSeriesDataBaseID timeseriesdata.TimeSeriesDataBaseID,
	data ...timeseriesdata.TimeSeriesData,
) error {
	if len(data) <= 0 {
		return nil
	}
	if err := t.Cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		for _, d := range data {
			docRef := t.firestoreDocTimeSeriesData(timeSeriesDataBaseID, d.GetID())
			if err := tx.Set(docRef, d); err != nil {
				return terrors.Wrap(err)
			}
		}
		return nil
	}); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
