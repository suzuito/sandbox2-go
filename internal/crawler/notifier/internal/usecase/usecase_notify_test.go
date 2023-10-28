package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
	"go.uber.org/mock/gomock"
)

type DummyTimeSeriesData struct {
	ID          string
	PublishedAt time.Time
}

func (t *DummyTimeSeriesData) GetID() timeseriesdata.TimeSeriesDataID {
	return timeseriesdata.TimeSeriesDataID(t.ID)
}
func (t *DummyTimeSeriesData) GetPublishedAt() time.Time {
	return t.PublishedAt
}

func toNotifiers(m []*notifier.MockNotifier) []notifier.Notifier {
	mm := []notifier.Notifier{}
	for _, v := range m {
		mm = append(mm, v)
	}
	return mm
}

func TestNotifyOnGCF(t *testing.T) {
	testCases := []struct {
		desc          string
		inputFullPath string
		expectedError string
		setUp         func(m *goMocks)
	}{
		{
			desc:          "3 notifiers",
			inputFullPath: "/hoge/fuga",
			setUp: func(m *goMocks) {
				notifiers := []*notifier.MockNotifier{
					notifier.NewMockNotifier(m.Controller),
					notifier.NewMockNotifier(m.Controller),
					notifier.NewMockNotifier(m.Controller),
				}
				notifiers[0].EXPECT().ID().Return(notifier.NotifierID("notifier1")).AnyTimes()
				notifiers[0].EXPECT().Notify(gomock.Any(), gomock.Any())
				notifiers[0].EXPECT().NewEmptyTimeSeriesData(gomock.Any()).Return(&DummyTimeSeriesData{})
				notifiers[1].EXPECT().ID().Return(notifier.NotifierID("notifier2")).AnyTimes()
				notifiers[1].EXPECT().Notify(gomock.Any(), gomock.Any())
				notifiers[1].EXPECT().NewEmptyTimeSeriesData(gomock.Any()).Return(&DummyTimeSeriesData{})
				notifiers[2].EXPECT().ID().Return(notifier.NotifierID("notifier3")).AnyTimes()
				notifiers[2].EXPECT().Notify(gomock.Any(), gomock.Any())
				notifiers[2].EXPECT().NewEmptyTimeSeriesData(gomock.Any()).Return(&DummyTimeSeriesData{})
				m.NotifierFactory.EXPECT().
					GetNotiferFromDocPathFirestore(
						gomock.Any(),
						"/hoge/fuga",
					).
					Return(toNotifiers(notifiers), nil)
				m.Repository.EXPECT().
					GetTimeSeriesDataFromFullPathFirestore(
						gomock.Any(),
						"/hoge/fuga",
						gomock.Any(),
					).Times(3)
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "NotifyOnGCP %s", gomock.Any()),
					m.L.EXPECT().Infof(gomock.Any(), "Notifier %s", notifier.NotifierID("notifier1")),
					m.L.EXPECT().Infof(gomock.Any(), "Notifier %s", notifier.NotifierID("notifier2")),
					m.L.EXPECT().Infof(gomock.Any(), "Notifier %s", notifier.NotifierID("notifier3")),
				)
			},
		},
		{
			desc:          "error from GetNotiferFromDocPathFirestore",
			inputFullPath: "/hoge/fuga",
			expectedError: "dummy error",
			setUp: func(m *goMocks) {
				m.NotifierFactory.EXPECT().
					GetNotiferFromDocPathFirestore(
						gomock.Any(),
						"/hoge/fuga",
					).
					Return(nil, errors.New("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "NotifyOnGCP %s", gomock.Any()),
				)
			},
		},
		{
			desc:          "3 notifiers and 1 error",
			inputFullPath: "/hoge/fuga",
			setUp: func(m *goMocks) {
				notifiers := []*notifier.MockNotifier{
					notifier.NewMockNotifier(m.Controller),
				}
				notifiers[0].EXPECT().ID().Return(notifier.NotifierID("notifier1")).AnyTimes()
				notifiers[0].EXPECT().NewEmptyTimeSeriesData(gomock.Any()).Return(&DummyTimeSeriesData{})
				m.NotifierFactory.EXPECT().
					GetNotiferFromDocPathFirestore(
						gomock.Any(),
						"/hoge/fuga",
					).
					Return(toNotifiers(notifiers), nil)
				m.Repository.EXPECT().
					GetTimeSeriesDataFromFullPathFirestore(
						gomock.Any(),
						"/hoge/fuga",
						gomock.Any(),
					).
					Return(errors.New("dummy error"))
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "NotifyOnGCP %s", gomock.Any()),
					m.L.EXPECT().Infof(gomock.Any(), "Notifier %s", notifier.NotifierID("notifier1")),
					m.L.EXPECT().Errorf(gomock.Any(), "%+v on notifier %s", gomock.Any(), notifier.NotifierID("notifier1")),
				)
			},
		},
		{
			desc:          "3 notifiers",
			inputFullPath: "/hoge/fuga",
			setUp: func(m *goMocks) {
				notifiers := []*notifier.MockNotifier{
					notifier.NewMockNotifier(m.Controller),
					notifier.NewMockNotifier(m.Controller),
					notifier.NewMockNotifier(m.Controller),
				}
				notifiers[0].EXPECT().ID().Return(notifier.NotifierID("notifier1")).AnyTimes()
				notifiers[0].EXPECT().Notify(gomock.Any(), gomock.Any())
				notifiers[0].EXPECT().NewEmptyTimeSeriesData(gomock.Any()).Return(&DummyTimeSeriesData{})
				notifiers[1].EXPECT().ID().Return(notifier.NotifierID("notifier2")).AnyTimes()
				notifiers[1].EXPECT().Notify(gomock.Any(), gomock.Any()).Return(errors.New("dummy error"))
				notifiers[1].EXPECT().NewEmptyTimeSeriesData(gomock.Any()).Return(&DummyTimeSeriesData{})
				notifiers[2].EXPECT().ID().Return(notifier.NotifierID("notifier3")).AnyTimes()
				notifiers[2].EXPECT().Notify(gomock.Any(), gomock.Any())
				notifiers[2].EXPECT().NewEmptyTimeSeriesData(gomock.Any()).Return(&DummyTimeSeriesData{})
				m.NotifierFactory.EXPECT().
					GetNotiferFromDocPathFirestore(
						gomock.Any(),
						"/hoge/fuga",
					).
					Return(toNotifiers(notifiers), nil)
				m.Repository.EXPECT().
					GetTimeSeriesDataFromFullPathFirestore(
						gomock.Any(),
						"/hoge/fuga",
						gomock.Any(),
					).Times(3)
				gomock.InOrder(
					m.L.EXPECT().Infof(gomock.Any(), "NotifyOnGCP %s", gomock.Any()),
					m.L.EXPECT().Infof(gomock.Any(), "Notifier %s", notifier.NotifierID("notifier1")),
					m.L.EXPECT().Infof(gomock.Any(), "Notifier %s", notifier.NotifierID("notifier2")),
					m.L.EXPECT().Errorf(gomock.Any(), "%+v on notifier %s", gomock.Any(), notifier.NotifierID("notifier2")),
					m.L.EXPECT().Infof(gomock.Any(), "Notifier %s", notifier.NotifierID("notifier3")),
				)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := newMocks(t)
			defer m.Finish()
			tC.setUp(m)
			u := m.NewUsecase()
			err := u.NotifyOnGCF(context.Background(), tC.inputFullPath)
			if err == nil {
			}
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}
