package factoryimpl

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type DummyParser struct{}

func (t *DummyParser) ID() crawler.ParserID { return "dummy" }
func (t *DummyParser) Do(_ context.Context, _ io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	return nil, nil
}

func TestParserFactory(t *testing.T) {
	var dummyParser crawler.Parser = &DummyParser{}
	testCases := []struct {
		desc          string
		inputDef      crawler.ParserDefinition
		inputNewFuncs []NewFuncParser
		expectedError string
		expected      crawler.Parser
	}{
		{
			desc: `ParserをFactoryから取得できるケース
			inputNewFuncs[0]()ではErrNoMatchedParserIDが返されたが
			inputNewFuncs[1]()でParserが取得できた
			`,
			inputDef: crawler.ParserDefinition{
				ID: "dummy_id_01",
			},
			inputNewFuncs: []NewFuncParser{
				func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
					return nil, factoryerror.ErrNoMatchedParserID
				},
				func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
					return dummyParser, nil
				},
			},
			expected: dummyParser,
		},
		{
			desc: `NewNotifierFuncが意図しないエラーを返すケース`,
			inputDef: crawler.ParserDefinition{
				ID: "dummy_id_01",
			},
			inputNewFuncs: []NewFuncParser{
				func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
					return nil, fmt.Errorf("dummy")
				},
			},
			expectedError: "dummy",
		},
		{
			desc: `全てのinputNewFuncsがErrNoMatchedParserIDを返すケース`,
			inputDef: crawler.ParserDefinition{
				ID: "dummy_id_01",
			},
			inputNewFuncs: []NewFuncParser{
				func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
					return nil, factoryerror.ErrNoMatchedParserID
				},
				func(_ *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
					return nil, factoryerror.ErrNoMatchedParserID
				},
			},
			expectedError: "Parser 'dummy_id_01' is not found in available list",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := ParserFactory{
				NewFuncs: tC.inputNewFuncs,
			}
			notifierInstance, err := f.Get(context.Background(), &tC.inputDef)
			test_helper.AssertError(t, tC.expectedError, err)
			assert.Equal(t, tC.expected, notifierInstance)
		})
	}
}
