package bmysql

import "github.com/suzuito/sandbox2-go/blog/entity"

type tableTag struct {
	ID entity.TagID
}

func (t *tableTag) TableName() string {
	return "tags"
}

func newTableTag(article *entity.Tag) *tableTag {
	return &tableTag{
		ID: article.ID,
	}
}
