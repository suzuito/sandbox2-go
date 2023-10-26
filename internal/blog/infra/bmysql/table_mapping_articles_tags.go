package bmysql

import (
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type tableMappingArticlesTags struct {
	ArticleID      entity.ArticleID `gorm:"primaryKey"`
	ArticleVersion int32            `gorm:"primaryKey"`
	TagID          entity.TagID
}

func (t *tableMappingArticlesTags) TableName() string {
	return "mapping_articles_tags"
}
