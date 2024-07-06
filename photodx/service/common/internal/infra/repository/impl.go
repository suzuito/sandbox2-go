package repository

import (
	"time"

	"gorm.io/gorm"
)

type Impl struct {
	GormDB *gorm.DB

	NowFunc func() time.Time
}
