package viewmodel

import "github.com/suzuito/sandbox2-go/blog2/internal/entity"

type PageAdminArticles struct {
	ComponentCommonHead ComponentCommonHead
	Articles            []*entity.Article
}
