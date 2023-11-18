// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/pkg/entity/crawler/publisher.go
//
// Generated by this command:
//
//	mockgen -source crawler/pkg/entity/crawler/publisher.go -package crawler -destination crawler/pkg/entity/crawler/publisher_mock.go
//
// Package crawler is a generated GoMock package.
package crawler

import (
	context "context"
	reflect "reflect"

	timeseriesdata "github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	gomock "go.uber.org/mock/gomock"
)

// MockPublisher is a mock of Publisher interface.
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher.
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance.
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockPublisher) Do(ctx context.Context, input CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, input}
	for _, a := range data {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockPublisherMockRecorder) Do(ctx, input any, data ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, input}, data...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockPublisher)(nil).Do), varargs...)
}

// ID mocks base method.
func (m *MockPublisher) ID() PublisherID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(PublisherID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockPublisherMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockPublisher)(nil).ID))
}
