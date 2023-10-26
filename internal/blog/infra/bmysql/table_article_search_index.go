package bmysql

import (
	"time"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type tableArticleSearchIndex struct {
	ArticleID             entity.ArticleID
	CurrentArticleVersion int32
	Tags                  string
	Date                  time.Time
}

func (t *tableArticleSearchIndex) TableName() string {
	return "articles_search_index"
}
