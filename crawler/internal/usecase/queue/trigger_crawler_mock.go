// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/internal/usecase/queue/trigger_crawler.go
//
// Generated by this command:
//
//	mockgen -source crawler/internal/usecase/queue/trigger_crawler.go -package queue -destination crawler/internal/usecase/queue/trigger_crawler_mock.go
//
// Package queue is a generated GoMock package.
package queue

import (
	context "context"
	reflect "reflect"

	crawler "github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	gomock "go.uber.org/mock/gomock"
)

// MockTriggerCrawlerQueue is a mock of TriggerCrawlerQueue interface.
type MockTriggerCrawlerQueue struct {
	ctrl     *gomock.Controller
	recorder *MockTriggerCrawlerQueueMockRecorder
}

// MockTriggerCrawlerQueueMockRecorder is the mock recorder for MockTriggerCrawlerQueue.
type MockTriggerCrawlerQueueMockRecorder struct {
	mock *MockTriggerCrawlerQueue
}

// NewMockTriggerCrawlerQueue creates a new mock instance.
func NewMockTriggerCrawlerQueue(ctrl *gomock.Controller) *MockTriggerCrawlerQueue {
	mock := &MockTriggerCrawlerQueue{ctrl: ctrl}
	mock.recorder = &MockTriggerCrawlerQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTriggerCrawlerQueue) EXPECT() *MockTriggerCrawlerQueueMockRecorder {
	return m.recorder
}

// PublishCrawlEvent mocks base method.
func (m *MockTriggerCrawlerQueue) PublishCrawlEvent(ctx context.Context, crawlerID crawler.CrawlerID, crawlerInputData crawler.CrawlerInputData, crawlFunctionID crawler.CrawlFunctionID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishCrawlEvent", ctx, crawlerID, crawlerInputData, crawlFunctionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishCrawlEvent indicates an expected call of PublishCrawlEvent.
func (mr *MockTriggerCrawlerQueueMockRecorder) PublishCrawlEvent(ctx, crawlerID, crawlerInputData, crawlFunctionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishCrawlEvent", reflect.TypeOf((*MockTriggerCrawlerQueue)(nil).PublishCrawlEvent), ctx, crawlerID, crawlerInputData, crawlFunctionID)
}

// PublishDispatchCrawlEvent mocks base method.
func (m *MockTriggerCrawlerQueue) PublishDispatchCrawlEvent(ctx context.Context, crawlerID crawler.CrawlerID, crawlerInputData crawler.CrawlerInputData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishDispatchCrawlEvent", ctx, crawlerID, crawlerInputData)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishDispatchCrawlEvent indicates an expected call of PublishDispatchCrawlEvent.
func (mr *MockTriggerCrawlerQueueMockRecorder) PublishDispatchCrawlEvent(ctx, crawlerID, crawlerInputData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishDispatchCrawlEvent", reflect.TypeOf((*MockTriggerCrawlerQueue)(nil).PublishDispatchCrawlEvent), ctx, crawlerID, crawlerInputData)
}

// RecieveCrawlEvent mocks base method.
func (m *MockTriggerCrawlerQueue) RecieveCrawlEvent(ctx context.Context, rawBytes []byte) (crawler.CrawlerID, crawler.CrawlerInputData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecieveCrawlEvent", ctx, rawBytes)
	ret0, _ := ret[0].(crawler.CrawlerID)
	ret1, _ := ret[1].(crawler.CrawlerInputData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecieveCrawlEvent indicates an expected call of RecieveCrawlEvent.
func (mr *MockTriggerCrawlerQueueMockRecorder) RecieveCrawlEvent(ctx, rawBytes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecieveCrawlEvent", reflect.TypeOf((*MockTriggerCrawlerQueue)(nil).RecieveCrawlEvent), ctx, rawBytes)
}