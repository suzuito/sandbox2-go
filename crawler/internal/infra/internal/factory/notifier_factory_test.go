package factory

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type DummyNotifier struct{}

func (t *DummyNotifier) ID() notifier.NotifierID { return "dummy" }
func (t *DummyNotifier) NewEmptyTimeSeriesData() timeseriesdata.TimeSeriesData {
	return nil
}
func (t *DummyNotifier) Notify(
	_ context.Context,
	_ timeseriesdata.TimeSeriesData,
) error {
	return nil
}

func TestNotifierFactory(t *testing.T) {
	var dummyNotifier1 notifier.Notifier = &DummyNotifier{}
	testCases := []struct {
		desc          string
		inputDef      notifier.NotifierDefinition
		inputNewFuncs []NewFuncNotifier
		expectedError string
		expected      notifier.Notifier
	}{
		{
			desc: `NotifierをFactoryから取得できるケース
			inputNewFuncs[0]()ではErrNoMatchedNotifierIDが返されたが
			inputNewFuncs[1]()でNotifierが取得できた
			`,
			inputDef: notifier.NotifierDefinition{
				ID: "dummy_notifier_id_01",
			},
			inputNewFuncs: []NewFuncNotifier{
				func(_ *notifier.NotifierDefinition, _ *NewFuncNotifierArgument) (notifier.Notifier, error) {
					return nil, ErrNoMatchedNotifierID
				},
				func(_ *notifier.NotifierDefinition, _ *NewFuncNotifierArgument) (notifier.Notifier, error) {
					return dummyNotifier1, nil
				},
			},
			expected: dummyNotifier1,
		},
		{
			desc: `NewNotifierFuncが意図しないエラーを返すケース`,
			inputDef: notifier.NotifierDefinition{
				ID: notifier.NotifierID("dummy_notifier_id_01"),
			},
			inputNewFuncs: []NewFuncNotifier{
				func(_ *notifier.NotifierDefinition, _ *NewFuncNotifierArgument) (notifier.Notifier, error) {
					return nil, fmt.Errorf("dummy")
				},
			},
			expectedError: "dummy",
		},
		{
			desc: `全てのinputNewFuncsがErrNoMatchedNotifierIDを返すケース`,
			inputDef: notifier.NotifierDefinition{
				ID: notifier.NotifierID("dummy_notifier_id_01"),
			},
			inputNewFuncs: []NewFuncNotifier{
				func(_ *notifier.NotifierDefinition, _ *NewFuncNotifierArgument) (notifier.Notifier, error) {
					return nil, ErrNoMatchedNotifierID
				},
				func(_ *notifier.NotifierDefinition, _ *NewFuncNotifierArgument) (notifier.Notifier, error) {
					return nil, ErrNoMatchedNotifierID
				},
			},
			expectedError: "Notifier 'dummy_notifier_id_01' is not found in available list",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f := NotifierFactory{
				NewFuncs: tC.inputNewFuncs,
			}
			notifierInstance, err := f.Get(context.Background(), &tC.inputDef)
			test_helper.AssertError(t, tC.expectedError, err)
			assert.Equal(t, tC.expected, notifierInstance)
		})
	}
}
