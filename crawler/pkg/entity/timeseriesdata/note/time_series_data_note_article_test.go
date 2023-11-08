package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

func TestTimeSeriesDataNoteArticle(t *testing.T) {
	a := TimeSeriesDataNoteArticle{
		URL:         "https://www.example.com",
		PublishedAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
	}
	assert.Equal(t, timeseriesdata.TimeSeriesDataID("https---www.example.com"), a.GetID())
	assert.Equal(t, time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC), a.GetPublishedAt())
}
