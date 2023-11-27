package factory

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type DummyPublisher struct{}

func (t *DummyPublisher) ID() crawler.PublisherID { return "dummy" }
func (t *DummyPublisher) Do(_ context.Context, _ crawler.CrawlerInputData, _ ...timeseriesdata.TimeSeriesData) error {
	return nil
}

func TestPublisherFactory(t *testing.T) {
	var dummyPublisher crawler.Publisher = &DummyPublisher{}
	testCases := []struct {
		desc          string
		inputDef      crawler.PublisherDefinition
		inputNewFuncs []NewFuncPublisher
		expectedError string
		expected      crawler.Publisher
	}{
		{
			desc: `PublisherをFactoryから取得できるケース
			inputNewFuncs[0]()ではErrNoMatchedPublisherIDが返されたが
			inputNewFuncs[1]()でPublisherが取得できた
			`,
			inputDef: crawler.PublisherDefinition{
				ID: "dummy_id_01",
			},
			inputNewFuncs: []NewFuncPublisher{
				func(_ *crawler.PublisherDefinition, _ *NewFuncPublisherArgument) (crawler.Publisher, error) {
					return nil, ErrNoMatchedPublisherID
				},
				func(_ *crawler.PublisherDefinition, _ *NewFuncPublisherArgument) (crawler.Publisher, error) {
					return dummyPublisher, nil
				},
			},
			expected: dummyPublisher,
		},
		{
			desc: `NewFuncが意図しないエラーを返すケース`,
			inputDef: crawler.PublisherDefinition{
				ID: "dummy_id_01",
			},
			inputNewFuncs: []NewFuncPublisher{
				func(_ *crawler.PublisherDefinition, _ *NewFuncPublisherArgument) (crawler.Publisher, error) {
					return nil, fmt.Errorf("dummy")
				},
			},
			expectedError: "dummy",
		},
		{
			desc: `全てのinputNewFuncsがErrNoMatchedPublisherIDを返すケース`,
			inputDef: crawler.PublisherDefinition{
				ID: "dummy_id_01",
			},
			inputNewFuncs: []NewFuncPublisher{
				func(_ *crawler.PublisherDefinition, _ *NewFuncPublisherArgument) (crawler.Publisher, error) {
					return nil, ErrNoMatchedPublisherID
				},
				func(_ *crawler.PublisherDefinition, _ *NewFuncPublisherArgument) (crawler.Publisher, error) {
					return nil, ErrNoMatchedPublisherID
				},
			},
			expectedError: "Publisher 'dummy_id_01' is not found in available list",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := PublisherFactory{
				NewFuncs: tC.inputNewFuncs,
			}
			notifierInstance, err := f.Get(context.Background(), &tC.inputDef)
			test_helper.AssertError(t, tC.expectedError, err)
			assert.Equal(t, tC.expected, notifierInstance)
		})
	}
}
