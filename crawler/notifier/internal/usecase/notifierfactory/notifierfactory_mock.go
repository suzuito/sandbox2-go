// Code generated by MockGen. DO NOT EDIT.
// Source: crawler/notifier/internal/usecase/notifierfactory/notifierfactory.go
//
// Generated by this command:
//
//	mockgen -source crawler/notifier/internal/usecase/notifierfactory/notifierfactory.go -package notifierfactory -destination crawler/notifier/internal/usecase/notifierfactory/notifierfactory_mock.go
//
// Package notifierfactory is a generated GoMock package.
package notifierfactory

import (
	context "context"
	reflect "reflect"

	notifier "github.com/suzuito/sandbox2-go/crawler/notifier/internal/entity/notifier"
	gomock "go.uber.org/mock/gomock"
)

// MockNotifierFactory is a mock of NotifierFactory interface.
type MockNotifierFactory struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierFactoryMockRecorder
}

// MockNotifierFactoryMockRecorder is the mock recorder for MockNotifierFactory.
type MockNotifierFactoryMockRecorder struct {
	mock *MockNotifierFactory
}

// NewMockNotifierFactory creates a new mock instance.
func NewMockNotifierFactory(ctrl *gomock.Controller) *MockNotifierFactory {
	mock := &MockNotifierFactory{ctrl: ctrl}
	mock.recorder = &MockNotifierFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifierFactory) EXPECT() *MockNotifierFactoryMockRecorder {
	return m.recorder
}

// GetNotiferFromDocPathFirestore mocks base method.
func (m *MockNotifierFactory) GetNotiferFromDocPathFirestore(ctx context.Context, fullPath string) ([]notifier.Notifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotiferFromDocPathFirestore", ctx, fullPath)
	ret0, _ := ret[0].([]notifier.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotiferFromDocPathFirestore indicates an expected call of GetNotiferFromDocPathFirestore.
func (mr *MockNotifierFactoryMockRecorder) GetNotiferFromDocPathFirestore(ctx, fullPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotiferFromDocPathFirestore", reflect.TypeOf((*MockNotifierFactory)(nil).GetNotiferFromDocPathFirestore), ctx, fullPath)
}
