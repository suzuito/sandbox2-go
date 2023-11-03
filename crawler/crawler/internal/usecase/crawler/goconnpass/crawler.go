package goconnpass

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/internal/constant"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

const CrawlerID crawler.CrawlerID = "goconnpass"

type Crawler struct {
	repository repository.Repository
	cliHTTP    *http.Client
}

func NewCrawler(repository repository.Repository) crawler.Crawler {
	return &Crawler{
		repository: repository,
		cliHTTP:    http.DefaultClient,
	}
}

func (t *Crawler) ID() crawler.CrawlerID {
	return CrawlerID
}

func (t *Crawler) Name() string {
	return string(CrawlerID)
}

func (t *Crawler) Fetch(ctx context.Context, w io.Writer) error {
	u, _ := url.Parse("https://connpass.com/api/v1/event/")
	q := u.Query()
	q.Add("keyword_or", "go言語")
	q.Add("keyword_or", "golang")
	q.Add("keyword_or", "gopher")
	d := time.Now()
	for i := 0; i < 30; i++ {
		q.Add("ymd", d.Add(time.Duration(i)*time.Hour*24).Format("20060102"))
	}
	q.Add("count", "100")
	u.RawQuery = q.Encode()
	res, err := t.cliHTTP.Get(u.String())
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return terrors.Wrapf("HTTP error is occured code=%d", res.StatusCode)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Crawler) Parse(ctx context.Context, r io.Reader) ([]timeseriesdata.TimeSeriesData, error) {
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

type connpassAPIResponseEvent struct {
	EventID     int64  `json:"event_id"`
	Title       string `json:"title"`
	Catch       string `json:"catch"`
	Description string `json:"description"`
	EventURL    string `json:"event_url"`
	StartedAt   string `json:"started_at"`
	EndedAt     string `json:"ended_at"`
	Place       string `json:"place"`
	Address     string `json:"address"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

type connpassAPIResponse struct {
	Events []connpassAPIResponseEvent `json:"events"`
}

func (t *Crawler) Publish(ctx context.Context, data ...timeseriesdata.TimeSeriesData) error {
	return terrors.Wrap(t.repository.SetTimeSeriesData(ctx, CrawlerID, data...))
}
