// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/repository_article_source.go
//
// Generated by this command:
//
//	mockgen -source usecase/repository_article_source.go -package usecase -destination usecase/repository_article_source_mock.go
//
// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/suzuito/sandbox2-go/internal/blog/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryArticleSource is a mock of RepositoryArticleSource interface.
type MockRepositoryArticleSource struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryArticleSourceMockRecorder
}

// MockRepositoryArticleSourceMockRecorder is the mock recorder for MockRepositoryArticleSource.
type MockRepositoryArticleSourceMockRecorder struct {
	mock *MockRepositoryArticleSource
}

// NewMockRepositoryArticleSource creates a new mock instance.
func NewMockRepositoryArticleSource(ctrl *gomock.Controller) *MockRepositoryArticleSource {
	mock := &MockRepositoryArticleSource{ctrl: ctrl}
	mock.recorder = &MockRepositoryArticleSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryArticleSource) EXPECT() *MockRepositoryArticleSourceMockRecorder {
	return m.recorder
}

// GetArticleSource mocks base method.
func (m *MockRepositoryArticleSource) GetArticleSource(ctx context.Context, articleSourceID entity.ArticleSourceID, version string) (*entity.ArticleSource, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleSource", ctx, articleSourceID, version)
	ret0, _ := ret[0].(*entity.ArticleSource)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetArticleSource indicates an expected call of GetArticleSource.
func (mr *MockRepositoryArticleSourceMockRecorder) GetArticleSource(ctx, articleSourceID, version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleSource", reflect.TypeOf((*MockRepositoryArticleSource)(nil).GetArticleSource), ctx, articleSourceID, version)
}

// GetArticleSources mocks base method.
func (m *MockRepositoryArticleSource) GetArticleSources(ctx context.Context, ref string, proc func(*entity.ArticleSource, []byte) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleSources", ctx, ref, proc)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticleSources indicates an expected call of GetArticleSources.
func (mr *MockRepositoryArticleSourceMockRecorder) GetArticleSources(ctx, ref, proc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleSources", reflect.TypeOf((*MockRepositoryArticleSource)(nil).GetArticleSources), ctx, ref, proc)
}

// GetBranches mocks base method.
func (m *MockRepositoryArticleSource) GetBranches(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranches", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranches indicates an expected call of GetBranches.
func (mr *MockRepositoryArticleSourceMockRecorder) GetBranches(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranches", reflect.TypeOf((*MockRepositoryArticleSource)(nil).GetBranches), ctx)
}

// GetVersions mocks base method.
func (m *MockRepositoryArticleSource) GetVersions(ctx context.Context, branch string, articleSourceID entity.ArticleSourceID) ([]entity.ArticleSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersions", ctx, branch, articleSourceID)
	ret0, _ := ret[0].([]entity.ArticleSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersions indicates an expected call of GetVersions.
func (mr *MockRepositoryArticleSourceMockRecorder) GetVersions(ctx, branch, articleSourceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersions", reflect.TypeOf((*MockRepositoryArticleSource)(nil).GetVersions), ctx, branch, articleSourceID)
}

// SearchArticleSources mocks base method.
func (m *MockRepositoryArticleSource) SearchArticleSources(ctx context.Context, queryString string, proc func(*entity.ArticleSource) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchArticleSources", ctx, queryString, proc)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchArticleSources indicates an expected call of SearchArticleSources.
func (mr *MockRepositoryArticleSourceMockRecorder) SearchArticleSources(ctx, queryString, proc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchArticleSources", reflect.TypeOf((*MockRepositoryArticleSource)(nil).SearchArticleSources), ctx, queryString, proc)
}
