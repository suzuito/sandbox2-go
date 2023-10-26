package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ArticleSourceID string

type ArticleSource struct {
	ID      ArticleSourceID `validate:"required"`
	Version string          `validate:"required"`
	Meta    ArticleSourceMeta
}

type ArticleSourceMeta struct {
	URL string
}

func (t *ArticleSourceMeta) Scan(src any) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Cannot scan %+v", src)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return err
	}
	return nil
}

func (t ArticleSourceMeta) Value() (driver.Value, error) {
	return json.Marshal(&t)
}
