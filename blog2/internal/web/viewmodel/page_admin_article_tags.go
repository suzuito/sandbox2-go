package viewmodel

import "github.com/suzuito/sandbox2-go/blog2/internal/entity"

type PageAdminArticleTags struct {
	ComponentCommonHead ComponentCommonHead
	Article             *entity.Article
	Tags                []*entity.Tag
	JsEnv               PageAdminArticleTagsJsEnv
}

type PageAdminArticleTagsJsEnv struct{}
