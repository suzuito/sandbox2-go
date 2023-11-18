package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type TimeSeriesDataRepository interface {
	SetTimeSeriesData(
		ctx context.Context,
		timeSeriesDataBaseID timeseriesdata.TimeSeriesDataBaseID,
		timeSeriesData ...timeseriesdata.TimeSeriesData,
	) error
	GetTimeSeriesDataFromFullPathFirestore(
		ctx context.Context,
		fulPath string,
		d timeseriesdata.TimeSeriesData,
	) error
}
