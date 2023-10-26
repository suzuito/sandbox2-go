// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/repository_article.go
//
// Generated by this command:
//
//	mockgen -source usecase/repository_article.go -package usecase -destination usecase/repository_article_mock.go
//
// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/suzuito/sandbox2-go/internal/blog/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryArticle is a mock of RepositoryArticle interface.
type MockRepositoryArticle struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryArticleMockRecorder
}

// MockRepositoryArticleMockRecorder is the mock recorder for MockRepositoryArticle.
type MockRepositoryArticleMockRecorder struct {
	mock *MockRepositoryArticle
}

// NewMockRepositoryArticle creates a new mock instance.
func NewMockRepositoryArticle(ctrl *gomock.Controller) *MockRepositoryArticle {
	mock := &MockRepositoryArticle{ctrl: ctrl}
	mock.recorder = &MockRepositoryArticleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryArticle) EXPECT() *MockRepositoryArticleMockRecorder {
	return m.recorder
}

// GetArticleByPrimaryKey mocks base method.
func (m *MockRepositoryArticle) GetArticleByPrimaryKey(ctx context.Context, primaryKey entity.ArticlePrimaryKey, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByPrimaryKey", ctx, primaryKey, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticleByPrimaryKey indicates an expected call of GetArticleByPrimaryKey.
func (mr *MockRepositoryArticleMockRecorder) GetArticleByPrimaryKey(ctx, primaryKey, article any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByPrimaryKey", reflect.TypeOf((*MockRepositoryArticle)(nil).GetArticleByPrimaryKey), ctx, primaryKey, article)
}

// GetArticlesByArticleSourceID mocks base method.
func (m *MockRepositoryArticle) GetArticlesByArticleSourceID(ctx context.Context, articleSourceID entity.ArticleSourceID, articles *[]entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticlesByArticleSourceID", ctx, articleSourceID, articles)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticlesByArticleSourceID indicates an expected call of GetArticlesByArticleSourceID.
func (mr *MockRepositoryArticleMockRecorder) GetArticlesByArticleSourceID(ctx, articleSourceID, articles any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticlesByArticleSourceID", reflect.TypeOf((*MockRepositoryArticle)(nil).GetArticlesByArticleSourceID), ctx, articleSourceID, articles)
}

// GetArticlesByID mocks base method.
func (m *MockRepositoryArticle) GetArticlesByID(ctx context.Context, articleID entity.ArticleID, articles *[]entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticlesByID", ctx, articleID, articles)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticlesByID indicates an expected call of GetArticlesByID.
func (mr *MockRepositoryArticleMockRecorder) GetArticlesByID(ctx, articleID, articles any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticlesByID", reflect.TypeOf((*MockRepositoryArticle)(nil).GetArticlesByID), ctx, articleID, articles)
}

// GetArticlesByPrimaryKey mocks base method.
func (m *MockRepositoryArticle) GetArticlesByPrimaryKey(ctx context.Context, primaryKeys []entity.ArticlePrimaryKey, sortField SearchArticlesQuerySortField, sortOrder SortOrder, articles *[]entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticlesByPrimaryKey", ctx, primaryKeys, sortField, sortOrder, articles)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticlesByPrimaryKey indicates an expected call of GetArticlesByPrimaryKey.
func (mr *MockRepositoryArticleMockRecorder) GetArticlesByPrimaryKey(ctx, primaryKeys, sortField, sortOrder, articles any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticlesByPrimaryKey", reflect.TypeOf((*MockRepositoryArticle)(nil).GetArticlesByPrimaryKey), ctx, primaryKeys, sortField, sortOrder, articles)
}

// GetLatestArticle mocks base method.
func (m *MockRepositoryArticle) GetLatestArticle(ctx context.Context, articleID entity.ArticleID, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestArticle", ctx, articleID, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetLatestArticle indicates an expected call of GetLatestArticle.
func (mr *MockRepositoryArticleMockRecorder) GetLatestArticle(ctx, articleID, article any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestArticle", reflect.TypeOf((*MockRepositoryArticle)(nil).GetLatestArticle), ctx, articleID, article)
}

// SearchArticles mocks base method.
func (m *MockRepositoryArticle) SearchArticles(ctx context.Context, query SearchArticlesQuery, articlePrimaryKeys *[]entity.ArticlePrimaryKey, hasNext *bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchArticles", ctx, query, articlePrimaryKeys, hasNext)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchArticles indicates an expected call of SearchArticles.
func (mr *MockRepositoryArticleMockRecorder) SearchArticles(ctx, query, articlePrimaryKeys, hasNext any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchArticles", reflect.TypeOf((*MockRepositoryArticle)(nil).SearchArticles), ctx, query, articlePrimaryKeys, hasNext)
}

// SetArticle mocks base method.
func (m *MockRepositoryArticle) SetArticle(ctx context.Context, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetArticle", ctx, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetArticle indicates an expected call of SetArticle.
func (mr *MockRepositoryArticleMockRecorder) SetArticle(ctx, article any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArticle", reflect.TypeOf((*MockRepositoryArticle)(nil).SetArticle), ctx, article)
}

// SetArticleSearchIndex mocks base method.
func (m *MockRepositoryArticle) SetArticleSearchIndex(ctx context.Context, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetArticleSearchIndex", ctx, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetArticleSearchIndex indicates an expected call of SetArticleSearchIndex.
func (mr *MockRepositoryArticleMockRecorder) SetArticleSearchIndex(ctx, article any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArticleSearchIndex", reflect.TypeOf((*MockRepositoryArticle)(nil).SetArticleSearchIndex), ctx, article)
}
