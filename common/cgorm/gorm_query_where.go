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

type ListQuery2KeyDirection string

const (
	ListQuery2KeyDirectionAsc  ListQuery2KeyDirection = "asc"
	ListQuery2KeyDirectionDesc ListQuery2KeyDirection = "desc"
)

type ListQuery2BoundType string

const (
	ListQuery2BoundTypeLower ListQuery2BoundType = "lower"
	ListQuery2BoundTypeUpper ListQuery2BoundType = "upper"
	ListQuery2BoundTypeOnly  ListQuery2BoundType = "only"
)

type ListQuery2Bound struct {
	Type      ListQuery2BoundType
	KeyName   string
	Direction ListQuery2KeyDirection
	Open      bool
	Value     any
}

type ListQuery2KeyRange struct {
	Bounds []ListQuery2Bound
}

type ListQuery2 struct {
	Range ListQuery2KeyRange
	Limit int
}

func (t *ListQuery2) Set(db *gorm.DB) *gorm.DB {
	for _, bound := range t.Range.Bounds {
		o := "="
		switch bound.Type {
		case ListQuery2BoundTypeLower:
			o = "<"
			if !bound.Open {
				o = o + "="
			}
		case ListQuery2BoundTypeUpper:
			o = ">"
			if !bound.Open {
				o = o + "="
			}
		case ListQuery2BoundTypeOnly:
			o = "="
		}
		db = db.Where(fmt.Sprintf("%s %s ?", bound.KeyName, o), bound.Value)
	}
	db = db.Limit(t.Limit)
	for _, bound := range t.Range.Bounds {
		db = db.Order(fmt.Sprintf("%s %s", bound.KeyName, bound.Direction))
	}
	return db
}
