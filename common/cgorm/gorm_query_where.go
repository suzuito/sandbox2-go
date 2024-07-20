package cgorm

import (
	"fmt"

	"gorm.io/gorm"
)

type GormQueryWhere struct {
	Query any
	Args  []any
}

type SortColumnType string

var (
	Asc  SortColumnType = "ASC"
	Desc SortColumnType = "DESC"
)

type SortColumn struct {
	Name string         `json:"name"`
	Type SortColumnType `json:"type"`
}

type ListQuery struct {
	Offset      int          `json:"offset"`
	Limit       int          `json:"limit"`
	SortColumns []SortColumn `json:"sortColumns"`
}

func (t *ListQuery) Set(db *gorm.DB) *gorm.DB {
	for _, sortColumn := range t.SortColumns {
		db = db.Order(
			fmt.Sprintf("%s %s", sortColumn.Name, sortColumn.Type),
		)
	}
	db = db.Offset(t.Offset)
	db = db.Limit(t.Limit)
	return db
}

func (t *ListQuery) NextOffset() int {
	return t.Offset + t.Limit
}

func (t *ListQuery) PrevOffset() int {
	offset := t.Offset - t.Limit
	if offset < 0 {
		offset = 0
	}
	return offset
}

func (t *ListQuery) HasPrev() bool {
	return t.PrevOffset() >= 0
}
