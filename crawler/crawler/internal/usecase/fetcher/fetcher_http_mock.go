// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/crawler/internal/usecase/fetcher/fetcher_http.go
//
// Generated by this command:
//
//	mockgen -source crawler/crawler/internal/usecase/fetcher/fetcher_http.go -package fetcher -destination crawler/crawler/internal/usecase/fetcher/fetcher_http_mock.go
//
// Package fetcher is a generated GoMock package.
package fetcher

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFetcherHTTP is a mock of FetcherHTTP interface.
type MockFetcherHTTP struct {
	ctrl     *gomock.Controller
	recorder *MockFetcherHTTPMockRecorder
}

// MockFetcherHTTPMockRecorder is the mock recorder for MockFetcherHTTP.
type MockFetcherHTTPMockRecorder struct {
	mock *MockFetcherHTTP
}

// NewMockFetcherHTTP creates a new mock instance.
func NewMockFetcherHTTP(ctrl *gomock.Controller) *MockFetcherHTTP {
	mock := &MockFetcherHTTP{ctrl: ctrl}
	mock.recorder = &MockFetcherHTTPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFetcherHTTP) EXPECT() *MockFetcherHTTPMockRecorder {
	return m.recorder
}

// DoRequest mocks base method.
func (m *MockFetcherHTTP) DoRequest(ctx context.Context, request *http.Request, w io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRequest", ctx, request, w)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoRequest indicates an expected call of DoRequest.
func (mr *MockFetcherHTTPMockRecorder) DoRequest(ctx, request, w any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRequest", reflect.TypeOf((*MockFetcherHTTP)(nil).DoRequest), ctx, request, w)
}
