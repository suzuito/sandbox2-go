package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

type Repository interface {
	GetTimeSeriesDataFromFullPathFirestore(
		ctx context.Context,
		fulPath string,
		d timeseriesdata.TimeSeriesData,
	) error
}
