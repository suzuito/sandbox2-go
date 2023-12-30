// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/internal/usecase/repository/notifier.go
//
// Generated by this command:
//
//	mockgen -source crawler/internal/usecase/repository/notifier.go -package repository -destination crawler/internal/usecase/repository/notifier_mock.go
//
// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	notifier "github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	gomock "go.uber.org/mock/gomock"
)

// MockNotifierRepository is a mock of NotifierRepository interface.
type MockNotifierRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierRepositoryMockRecorder
}

// MockNotifierRepositoryMockRecorder is the mock recorder for MockNotifierRepository.
type MockNotifierRepositoryMockRecorder struct {
	mock *MockNotifierRepository
}

// NewMockNotifierRepository creates a new mock instance.
func NewMockNotifierRepository(ctrl *gomock.Controller) *MockNotifierRepository {
	mock := &MockNotifierRepository{ctrl: ctrl}
	mock.recorder = &MockNotifierRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifierRepository) EXPECT() *MockNotifierRepositoryMockRecorder {
	return m.recorder
}

// GetNotiferDefinitionsFromDocPathFirestore mocks base method.
func (m *MockNotifierRepository) GetNotiferDefinitionsFromDocPathFirestore(ctx context.Context, fullPath string) ([]notifier.NotifierDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotiferDefinitionsFromDocPathFirestore", ctx, fullPath)
	ret0, _ := ret[0].([]notifier.NotifierDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotiferDefinitionsFromDocPathFirestore indicates an expected call of GetNotiferDefinitionsFromDocPathFirestore.
func (mr *MockNotifierRepositoryMockRecorder) GetNotiferDefinitionsFromDocPathFirestore(ctx, fullPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotiferDefinitionsFromDocPathFirestore", reflect.TypeOf((*MockNotifierRepository)(nil).GetNotiferDefinitionsFromDocPathFirestore), ctx, fullPath)
}
