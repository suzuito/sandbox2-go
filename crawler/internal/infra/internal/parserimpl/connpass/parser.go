package connpass

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/constant"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Parser struct {
}

func (t *Parser) ID() crawler.ParserID {
	return "connpass"
}

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
		returned = append(returned, &timeseriesdata.TimeSeriesDataEvent{
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
			Organizer: &timeseriesdata.TimeSeriesDataEventOrganizer{
				Name:     event.OwnerDisplayName,
				URL:      fmt.Sprintf("https://connpass.com/%s", event.OwnerNickName),
				ImageURL: "https://connpass.com/static/img/api/connpass_logo_4.png",
			},
		})
	}
	return returned, nil
}

func New(def *crawler.ParserDefinition, _ *factory.NewFuncParserArgument) (crawler.Parser, error) {
	parser := Parser{}
	if def.ID != parser.ID() {
		return nil, factory.ErrNoMatchedParserID
	}
	return &parser, nil
}

type connpassAPIResponseEvent struct {
	EventID          int64  `json:"event_id"`
	Title            string `json:"title"`
	Catch            string `json:"catch"`
	Description      string `json:"description"`
	EventURL         string `json:"event_url"`
	StartedAt        string `json:"started_at"`
	EndedAt          string `json:"ended_at"`
	Place            string `json:"place"`
	Address          string `json:"address"`
	Lat              string `json:"lat"`
	Lon              string `json:"lon"`
	OwnerNickName    string `json:"owner_nickname"`
	OwnerDisplayName string `json:"owner_display_name"`
}

type connpassAPIResponse struct {
	Events []connpassAPIResponseEvent `json:"events"`
}
