// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/internal/entity/crawler/parser.go
//
// Generated by this command:
//
//	mockgen -source crawler/internal/entity/crawler/parser.go -package crawler -destination crawler/internal/entity/crawler/parser_mock.go
//
// Package crawler is a generated GoMock package.
package crawler

import (
	context "context"
	io "io"
	reflect "reflect"

	timeseriesdata "github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
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

// Parse mocks base method.
func (m *MockParser) Parse(ctx context.Context, r io.Reader) ([]timeseriesdata.TimeSeriesData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", ctx, r)
	ret0, _ := ret[0].([]timeseriesdata.TimeSeriesData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockParserMockRecorder) Parse(ctx, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), ctx, r)
}