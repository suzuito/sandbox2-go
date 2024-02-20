package usecase

import "github.com/suzuito/sandbox2-go/blog2/internal/repository"

type Impl struct {
	RepositoryArticleIndex repository.RepositoryArticleIndex
	RepositoryArticle      repository.RepositoryArticle
}
