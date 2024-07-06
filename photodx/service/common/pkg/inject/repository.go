package inject

import (
	"time"

	internal_repository "github.com/suzuito/sandbox2-go/photodx/service/common/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

func NewRepository(
	gormDB *gorm.DB,
	nowFunc func() time.Time,
) repository.Repository {
	return &internal_repository.Impl{
		GormDB:  gormDB,
		NowFunc: nowFunc,
	}
}
