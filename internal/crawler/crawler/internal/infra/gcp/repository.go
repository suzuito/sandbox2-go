package gcp

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

type Repository struct {
	fcli                    *firestore.Client
	firestoreBaseCollection string
}

func NewRepository(
	fcli *firestore.Client,
	firestoreBaseCollection string,
) *Repository {
	return &Repository{
		fcli:                    fcli,
		firestoreBaseCollection: firestoreBaseCollection,
	}
}

func (t *Repository) SetTimeSeriesData(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
	data ...timeseriesdata.TimeSeriesData,
) error {
	if len(data) <= 0 {
		return nil
	}
	if err := t.fcli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		for _, d := range data {
			docRef := t.firestoreDocTimeSeriesData(crawlerID, d.GetID())
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
