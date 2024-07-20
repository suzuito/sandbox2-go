package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/businesslogic"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
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
		NowFunc:         time.Now,
		UserIDGenerator: &proc.IDGeneratorImpl{},
	}
}
