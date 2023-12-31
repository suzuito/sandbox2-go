// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/pkg/entity/crawler/parser.go
//
// Generated by this command:
//
//	mockgen -source crawler/pkg/entity/crawler/parser.go -package crawler -destination crawler/pkg/entity/crawler/parser_mock.go
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
func (m *MockParser) Do(ctx context.Context, r io.Reader, input CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
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

// ID mocks base method.
func (m *MockParser) ID() ParserID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(ParserID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockParserMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockParser)(nil).ID))
}
