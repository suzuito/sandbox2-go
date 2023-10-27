// Code generated by MockGen. DO NOT EDIT.
// Source: internal/crawler/crawler/internal/entity/crawler/crawler.go
//
// Generated by this command:
//
//	mockgen -source internal/crawler/crawler/internal/entity/crawler/crawler.go -package crawler -destination internal/crawler/crawler/internal/entity/crawler/crawler_mock.go
//
// Package crawler is a generated GoMock package.
package crawler

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCrawler is a mock of Crawler interface.
type MockCrawler struct {
	ctrl     *gomock.Controller
	recorder *MockCrawlerMockRecorder
}

// MockCrawlerMockRecorder is the mock recorder for MockCrawler.
type MockCrawlerMockRecorder struct {
	mock *MockCrawler
}

// NewMockCrawler creates a new mock instance.
func NewMockCrawler(ctrl *gomock.Controller) *MockCrawler {
	mock := &MockCrawler{ctrl: ctrl}
	mock.recorder = &MockCrawlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrawler) EXPECT() *MockCrawlerMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockCrawler) ID() CrawlerID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(CrawlerID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockCrawlerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockCrawler)(nil).ID))
}

// Name mocks base method.
func (m *MockCrawler) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockCrawlerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCrawler)(nil).Name))
}

// NewFetcher mocks base method.
func (m *MockCrawler) NewFetcher(ctx context.Context) (Fetcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFetcher", ctx)
	ret0, _ := ret[0].(Fetcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFetcher indicates an expected call of NewFetcher.
func (mr *MockCrawlerMockRecorder) NewFetcher(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFetcher", reflect.TypeOf((*MockCrawler)(nil).NewFetcher), ctx)
}

// NewParser mocks base method.
func (m *MockCrawler) NewParser(ctx context.Context) (Parser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewParser", ctx)
	ret0, _ := ret[0].(Parser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewParser indicates an expected call of NewParser.
func (mr *MockCrawlerMockRecorder) NewParser(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewParser", reflect.TypeOf((*MockCrawler)(nil).NewParser), ctx)
}

// NewPublisher mocks base method.
func (m *MockCrawler) NewPublisher(ctx context.Context) (Publisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPublisher", ctx)
	ret0, _ := ret[0].(Publisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPublisher indicates an expected call of NewPublisher.
func (mr *MockCrawlerMockRecorder) NewPublisher(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPublisher", reflect.TypeOf((*MockCrawler)(nil).NewPublisher), ctx)
}
