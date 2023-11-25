package factory

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type DummyFetcher struct{}

func (t *DummyFetcher) ID() crawler.FetcherID { return "dummy" }
func (t *DummyFetcher) Do(_ context.Context, _ io.Writer, _ crawler.CrawlerInputData) error {
	return nil
}

func Test(t *testing.T) {
	var dummyFetcher1 crawler.Fetcher = &DummyFetcher{}
	testCases := []struct {
		desc            string
		inputDef        crawler.FetcherDefinition
		inputNewFuncs   []NewFuncFetcher
		expectedError   string
		expectedFetcher crawler.Fetcher
	}{
		{
			desc: `FetcherをFactoryから取得できるケース
			inputNewFuncs[0]()ではErrNoMatchedFetcherIDが返されたが
			inputNewFuncs[1]()でFetcherが取得できた
			`,
			inputDef: crawler.FetcherDefinition{
				ID: crawler.FetcherID("dummy_fetcher_id_01"),
			},
			inputNewFuncs: []NewFuncFetcher{
				func(_ *crawler.FetcherDefinition, _ *NewFuncFetcherArgument) (crawler.Fetcher, error) {
					return nil, ErrNoMatchedFetcherID
				},
				func(_ *crawler.FetcherDefinition, _ *NewFuncFetcherArgument) (crawler.Fetcher, error) {
					return dummyFetcher1, nil
				},
			},
			expectedFetcher: dummyFetcher1,
		},
		{
			desc: `NewFetcherFuncが意図しないエラーを返すケース`,
			inputDef: crawler.FetcherDefinition{
				ID: crawler.FetcherID("dummy_fetcher_id_01"),
			},
			inputNewFuncs: []NewFuncFetcher{
				func(_ *crawler.FetcherDefinition, _ *NewFuncFetcherArgument) (crawler.Fetcher, error) {
					return nil, fmt.Errorf("dummy")
				},
			},
			expectedError: "dummy",
		},
		{
			desc: `全てのinputNewFuncsがErrNoMatchedFetcherIDを返すケース`,
			inputDef: crawler.FetcherDefinition{
				ID: crawler.FetcherID("dummy_fetcher_id_01"),
			},
			inputNewFuncs: []NewFuncFetcher{
				func(_ *crawler.FetcherDefinition, _ *NewFuncFetcherArgument) (crawler.Fetcher, error) {
					return nil, ErrNoMatchedFetcherID
				},
				func(_ *crawler.FetcherDefinition, _ *NewFuncFetcherArgument) (crawler.Fetcher, error) {
					return nil, ErrNoMatchedFetcherID
				},
			},
			expectedError: "Fetcher 'dummy_fetcher_id_01' is not found in available list",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := FetcherFactory{
				HTTPClient: http.DefaultClient,
				NewFuncs:   tC.inputNewFuncs,
			}
			fetcher, err := f.Get(context.Background(), &tC.inputDef)
			test_helper.AssertError(t, tC.expectedError, err)
			assert.Equal(t, tC.expectedFetcher, fetcher)
		})
	}
}
