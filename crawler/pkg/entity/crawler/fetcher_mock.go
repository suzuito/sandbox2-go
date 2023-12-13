// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/pkg/entity/crawler/fetcher.go
//
// Generated by this command:
//
//	mockgen -source crawler/pkg/entity/crawler/fetcher.go -package crawler -destination crawler/pkg/entity/crawler/fetcher_mock.go
//
// Package crawler is a generated GoMock package.
package crawler

import (
	context "context"
	io "io"
	slog "log/slog"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFetcher is a mock of Fetcher interface.
type MockFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockFetcherMockRecorder
}

// MockFetcherMockRecorder is the mock recorder for MockFetcher.
type MockFetcherMockRecorder struct {
	mock *MockFetcher
}

// NewMockFetcher creates a new mock instance.
func NewMockFetcher(ctrl *gomock.Controller) *MockFetcher {
	mock := &MockFetcher{ctrl: ctrl}
	mock.recorder = &MockFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFetcher) EXPECT() *MockFetcherMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockFetcher) Do(ctx context.Context, logger *slog.Logger, w io.Writer, input CrawlerInputData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", ctx, logger, w, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockFetcherMockRecorder) Do(ctx, logger, w, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockFetcher)(nil).Do), ctx, logger, w, input)
}

// ID mocks base method.
func (m *MockFetcher) ID() FetcherID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(FetcherID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockFetcherMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockFetcher)(nil).ID))
}
