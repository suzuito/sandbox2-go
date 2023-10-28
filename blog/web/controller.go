package web

import (
	"github.com/suzuito/sandbox2-go/blog/usecase"
	"github.com/suzuito/sandbox2-go/blog/usecase/markdown2html"
)

type ControllerImpl struct {
	UC                      usecase.Usecase
	RepositoryArticle       usecase.RepositoryArticle
	RepositoryArticleSource usecase.RepositoryArticleSource
	RepositoryArticleHTML   usecase.RepositoryArticleHTML
	Markdown2HTML           markdown2html.Markdown2HTML
	Presenters              Presenters
	Setting                 *WebSetting
}

func (t *ControllerImpl) NoIndex() bool {
	return t.Setting.NoIndex
}

func (t *ControllerImpl) DirPathTemplates() string {
	return t.Setting.DirPathTemplates
}

func (t *ControllerImpl) DirPathCSS() string {
	return t.Setting.DirPathCSS
}

func (t *ControllerImpl) DirPathImages() string {
	return t.Setting.DirPathImages
}
