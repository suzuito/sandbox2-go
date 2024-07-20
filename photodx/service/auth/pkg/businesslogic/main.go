package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/repository"
	"gorm.io/gorm"
)

func Main(
	gormDB *gorm.DB,
) ExposedBusinessLogic {
	return &businesslogic.Impl{
		Repository: &infra_repository.Impl{
			GormDB:  gormDB,
			NowFunc: time.Now,
		},
		NowFunc: time.Now,
	}
}
