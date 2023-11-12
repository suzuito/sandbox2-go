package goconnpass

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/internal/constant"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Parser struct{}

func (t *Parser) Do(ctx context.Context, r io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	res := connpassAPIResponse{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	for _, event := range res.Events {
		startedAt, err := time.Parse(time.RFC3339, event.StartedAt)
		if err != nil {
			startedAt = time.Date(0, 0, 0, 0, 0, 0, 0, nil)
		}
		endedAt, err := time.Parse(time.RFC3339, event.EndedAt)
		if err != nil {
			endedAt = time.Date(0, 0, 0, 0, 0, 0, 0, nil)
		}
		startedAt = startedAt.In(constant.JST)
		endedAt = endedAt.In(constant.JST)
		lat, err := strconv.ParseFloat(event.Lat, 64)
		if err != nil {
			lat = -1
		}
		lon, err := strconv.ParseFloat(event.Lon, 64)
		if err != nil {
			lon = -1
		}
		description := ""
		doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(event.Description))
		if err != nil {
			description = ""
		} else {
			description = doc.Text()
		}
		returned = append(returned, &timeseriesdata.TimeSeriesDataConnpassEvent{
			EventID:     event.EventID,
			Title:       event.Title,
			Catch:       event.Catch,
			Description: description,
			EventURL:    event.EventURL,
			StartedAt:   startedAt,
			EndedAt:     endedAt,
			Place:       event.Place,
			Address:     event.Address,
			Lat:         lat,
			Lon:         lon,
		})
	}
	return returned, nil
}

func NewParser() crawler.Parser {
	return &Parser{}
}
