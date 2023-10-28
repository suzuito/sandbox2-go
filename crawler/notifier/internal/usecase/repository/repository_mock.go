// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/notifier/internal/usecase/repository/repository.go
//
// Generated by this command:
//
//	mockgen -source crawler/notifier/internal/usecase/repository/repository.go -package repository -destination crawler/notifier/internal/usecase/repository/repository_mock.go
//
// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	timeseriesdata "github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetTimeSeriesDataFromFullPathFirestore mocks base method.
func (m *MockRepository) GetTimeSeriesDataFromFullPathFirestore(ctx context.Context, fulPath string, d timeseriesdata.TimeSeriesData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeSeriesDataFromFullPathFirestore", ctx, fulPath, d)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTimeSeriesDataFromFullPathFirestore indicates an expected call of GetTimeSeriesDataFromFullPathFirestore.
func (mr *MockRepositoryMockRecorder) GetTimeSeriesDataFromFullPathFirestore(ctx, fulPath, d any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeSeriesDataFromFullPathFirestore", reflect.TypeOf((*MockRepository)(nil).GetTimeSeriesDataFromFullPathFirestore), ctx, fulPath, d)
}
