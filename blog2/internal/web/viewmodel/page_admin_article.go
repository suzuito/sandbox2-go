package viewmodel

import "github.com/suzuito/sandbox2-go/blog2/internal/entity"

type PageAdminArticle struct {
	ComponentCommonHead ComponentCommonHead
	Article             *entity.Article
}
