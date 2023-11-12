// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/crawler/internal/entity/crawler/crawler2.go
//
// Generated by this command:
//
//	mockgen -source crawler/crawler/internal/entity/crawler/crawler2.go -package crawler -destination crawler/crawler/internal/entity/crawler/crawler2_mock.go
//
// Package crawler is a generated GoMock package.
package crawler

import (
	context "context"
	io "io"
	reflect "reflect"

	timeseriesdata "github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
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
func (m *MockFetcher) Do(ctx context.Context, w io.Writer, input CrawlerInputData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", ctx, w, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockFetcherMockRecorder) Do(ctx, w, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockFetcher)(nil).Do), ctx, w, input)
}

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser.
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockParser) Do(ctx context.Context, r io.Writer, input CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", ctx, r, input)
	ret0, _ := ret[0].([]timeseriesdata.TimeSeriesData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockParserMockRecorder) Do(ctx, r, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockParser)(nil).Do), ctx, r, input)
}

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
