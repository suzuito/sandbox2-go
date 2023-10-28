package gcp

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

func (t *Repository) firestoreDocTimeSeriesData(
	crawlerID crawler.CrawlerID,
	id timeseriesdata.TimeSeriesDataID,
) *firestore.DocumentRef {
	return t.fcli.Doc(fmt.Sprintf("%s/TimeSeriesData/%s/%s", t.firestoreBaseCollection, crawlerID, id))
}
