package bmysql

import "github.com/suzuito/sandbox2-go/internal/blog/entity"

type tableMappingArticlesSourceArticles struct {
	ArticleID            entity.ArticleID
	ArticleVersion       int32
	ArticleSourceID      entity.ArticleSourceID
	ArticleSourceVersion string
	Meta                 entity.ArticleSourceMeta
}

func (t *tableMappingArticlesSourceArticles) TableName() string {
	return "mapping_articles_source_articles"
}
