package repository

import (
	"math"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"gorm.io/gorm"
)

type Impl struct {
	NowFunc func() time.Time
	GormDB  *gorm.DB
}

func int64ToInt(v int64) (int, error) {
	if v < math.MinInt || v > math.MaxInt {
		return -1, terrors.Wrapf("convert error %d", v)
	}
	return int(v), nil
}
