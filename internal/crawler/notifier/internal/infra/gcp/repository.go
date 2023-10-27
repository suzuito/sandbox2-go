package gcp

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/internal/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

type Repository struct {
	fcli *firestore.Client
}

func (t *Repository) GetTimeSeriesDataFromFullPathFirestore(
	ctx context.Context,
	fullPath string,
	d timeseriesdata.TimeSeriesData,
) error {
	doc := t.fcli.Doc(fullPath)
	snp, err := doc.Get(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	if err := snp.DataTo(d); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func NewRepository(fcli *firestore.Client) *Repository {
	return &Repository{
		fcli: fcli,
	}
}
