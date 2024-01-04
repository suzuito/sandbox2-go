package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"go.uber.org/mock/gomock"
)

func TestNotifyOnGCF(t *testing.T) {
	testCases := []struct {
		utTestCase
		inputFullPath string
		expectedError string
	}{
		{
			utTestCase: utTestCase{
				desc: "Success",
				setUp: func(mocks *utMocks) {
					mocks.MockNotifierRepository.EXPECT().
						GetNotiferDefinitionsFromDocPathFirestore(
							gomock.Any(),
							"dummy",
						).
						Return(
							[]notifier.NotifierDefinition{
								{ID: "notifier001"},
								{ID: "notifier002"},
							},
							nil,
						)
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier001"}).
						Return(mocks.MockNotifier, nil)
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier002"}).
						Return(mocks.MockNotifier, nil)
					mocks.MockNotifier.EXPECT().
						NewEmptyTimeSeriesData().
						Return(&timeseriesdata.TimeSeriesDataEvent{}).Times(2)
					mocks.MockTimeSeriesDataRepository.EXPECT().
						GetTimeSeriesDataFromFullPathFirestore(
							gomock.Any(),
							"dummy",
							gomock.Any(),
						).Return(nil).Times(2)
					mocks.MockNotifier.EXPECT().
						Notify(
							gomock.Any(),
							gomock.Any(),
						).Return(nil).Times(2)
				},
				expectedLogLines: []string{
					"level=INFO msg=NotifyOnGCP fullPath=dummy",
					"level=INFO msg=Notifier notifierID=notifier001",
					"level=INFO msg=Notifier notifierID=notifier002",
				},
			},
			inputFullPath: "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to GetNotiferDefinitionsFromDocPathFirestore",
				setUp: func(mocks *utMocks) {
					mocks.MockNotifierRepository.EXPECT().
						GetNotiferDefinitionsFromDocPathFirestore(
							gomock.Any(),
							"dummy",
						).
						Return(
							nil,
							errors.New("dummy"),
						)
				},
				expectedLogLines: []string{
					"level=INFO msg=NotifyOnGCP fullPath=dummy",
				},
			},
			inputFullPath: "dummy",
			expectedError: "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to NotifierFactory.Get",
				setUp: func(mocks *utMocks) {
					mocks.MockNotifierRepository.EXPECT().
						GetNotiferDefinitionsFromDocPathFirestore(
							gomock.Any(),
							"dummy",
						).
						Return(
							[]notifier.NotifierDefinition{
								{ID: "notifier001"},
								{ID: "notifier002"},
							},
							nil,
						)
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier001"}).
						Return(nil, errors.New("dummy"))
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier002"}).
						Return(mocks.MockNotifier, nil)
					mocks.MockNotifier.EXPECT().
						NewEmptyTimeSeriesData().
						Return(&timeseriesdata.TimeSeriesDataEvent{}).Times(1)
					mocks.MockTimeSeriesDataRepository.EXPECT().
						GetTimeSeriesDataFromFullPathFirestore(
							gomock.Any(),
							"dummy",
							gomock.Any(),
						).Return(nil).Times(1)
					mocks.MockNotifier.EXPECT().
						Notify(
							gomock.Any(),
							gomock.Any(),
						).Return(nil).Times(1)
				},
				expectedLogLines: []string{
					"level=INFO msg=NotifyOnGCP fullPath=dummy",
					"level=INFO msg=Notifier notifierID=notifier001",
					`level=ERROR msg="Failed to NotifierFactory.Get" notifierID=notifier001 err=dummy`,
					"level=INFO msg=Notifier notifierID=notifier002",
				},
			},
			inputFullPath: "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Failed to GetTimeSeriesDataFromFullPathFirestore",
				setUp: func(mocks *utMocks) {
					mocks.MockNotifierRepository.EXPECT().
						GetNotiferDefinitionsFromDocPathFirestore(
							gomock.Any(),
							"dummy",
						).
						Return(
							[]notifier.NotifierDefinition{
								{ID: "notifier001"},
								{ID: "notifier002"},
							},
							nil,
						)
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier001"}).
						Return(mocks.MockNotifier, nil)
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier002"}).
						Return(mocks.MockNotifier, nil)
					mocks.MockNotifier.EXPECT().
						NewEmptyTimeSeriesData().
						Return(&timeseriesdata.TimeSeriesDataEvent{}).Times(2)
					gomock.InOrder(
						mocks.MockTimeSeriesDataRepository.EXPECT().
							GetTimeSeriesDataFromFullPathFirestore(
								gomock.Any(),
								"dummy",
								gomock.Any(),
							).Return(errors.New("dummy")).Times(1),
						mocks.MockTimeSeriesDataRepository.EXPECT().
							GetTimeSeriesDataFromFullPathFirestore(
								gomock.Any(),
								"dummy",
								gomock.Any(),
							).Return(nil).Times(1),
					)
					mocks.MockNotifier.EXPECT().
						Notify(
							gomock.Any(),
							gomock.Any(),
						).Return(nil).Times(1)
				},
				expectedLogLines: []string{
					"level=INFO msg=NotifyOnGCP fullPath=dummy",
					"level=INFO msg=Notifier notifierID=notifier001",
					`level=ERROR msg="Failed to GetTimeSeriesDataFromFullPathFirestore" err=dummy`,
					"level=INFO msg=Notifier notifierID=notifier002",
				},
			},
			inputFullPath: "dummy",
		},
		{
			utTestCase: utTestCase{
				desc: "Success",
				setUp: func(mocks *utMocks) {
					mocks.MockNotifierRepository.EXPECT().
						GetNotiferDefinitionsFromDocPathFirestore(
							gomock.Any(),
							"dummy",
						).
						Return(
							[]notifier.NotifierDefinition{
								{ID: "notifier001"},
								{ID: "notifier002"},
							},
							nil,
						)
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier001"}).
						Return(mocks.MockNotifier, nil)
					mocks.MockNotifierFactory.EXPECT().
						Get(gomock.Any(), &notifier.NotifierDefinition{ID: "notifier002"}).
						Return(mocks.MockNotifier, nil)
					mocks.MockNotifier.EXPECT().
						NewEmptyTimeSeriesData().
						Return(&timeseriesdata.TimeSeriesDataEvent{}).Times(2)
					mocks.MockTimeSeriesDataRepository.EXPECT().
						GetTimeSeriesDataFromFullPathFirestore(
							gomock.Any(),
							"dummy",
							gomock.Any(),
						).Return(nil).Times(2)
					gomock.InOrder(
						mocks.MockNotifier.EXPECT().
							Notify(
								gomock.Any(),
								gomock.Any(),
							).Return(errors.New("dummy")).Times(1),
						mocks.MockNotifier.EXPECT().
							Notify(
								gomock.Any(),
								gomock.Any(),
							).Return(nil).Times(1),
					)
				},
				expectedLogLines: []string{
					"level=INFO msg=NotifyOnGCP fullPath=dummy",
					"level=INFO msg=Notifier notifierID=notifier001",
					`level=ERROR msg="Failed to Notify" err=dummy`,
					"level=INFO msg=Notifier notifierID=notifier002",
				},
			},
			inputFullPath: "dummy",
		},
	}
	for _, tC := range testCases {
		tC.run(t, func(uc *UsecaseImpl) {
			err := uc.NotifyOnGCF(context.Background(), tC.inputFullPath)
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
