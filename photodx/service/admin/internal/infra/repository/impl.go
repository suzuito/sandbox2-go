package repository

import (
	"time"

	"gorm.io/gorm"
)

type Impl struct {
	NowFunc func() time.Time
	GormDB  *gorm.DB
}
