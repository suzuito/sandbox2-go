// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/crawler/internal/usecase/queue/queue.go
//
// Generated by this command:
//
//	mockgen -source crawler/crawler/internal/usecase/queue/queue.go -package queue -destination crawler/crawler/internal/usecase/queue/queue_mock.go
//
// Package queue is a generated GoMock package.
package queue

import (
	context "context"
	reflect "reflect"

	crawler "github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	gomock "go.uber.org/mock/gomock"
)

// MockQueue is a mock of Queue interface.
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *MockQueueMockRecorder
}

// MockQueueMockRecorder is the mock recorder for MockQueue.
type MockQueueMockRecorder struct {
	mock *MockQueue
}

// NewMockQueue creates a new mock instance.
func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &MockQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueue) EXPECT() *MockQueueMockRecorder {
	return m.recorder
}

// PublishCrawlEvent mocks base method.
func (m *MockQueue) PublishCrawlEvent(ctx context.Context, crawlerID crawler.CrawlerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishCrawlEvent", ctx, crawlerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishCrawlEvent indicates an expected call of PublishCrawlEvent.
func (mr *MockQueueMockRecorder) PublishCrawlEvent(ctx, crawlerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishCrawlEvent", reflect.TypeOf((*MockQueue)(nil).PublishCrawlEvent), ctx, crawlerID)
}

// RecieveCrawlEvent mocks base method.
func (m *MockQueue) RecieveCrawlEvent(ctx context.Context, rawBytes []byte) (crawler.CrawlerID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecieveCrawlEvent", ctx, rawBytes)
	ret0, _ := ret[0].(crawler.CrawlerID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecieveCrawlEvent indicates an expected call of RecieveCrawlEvent.
func (mr *MockQueueMockRecorder) RecieveCrawlEvent(ctx, rawBytes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecieveCrawlEvent", reflect.TypeOf((*MockQueue)(nil).RecieveCrawlEvent), ctx, rawBytes)
}
