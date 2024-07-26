package businesslogic

import (
	"log/slog"
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/admin/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"gorm.io/gorm"
)

func Main(
	l *slog.Logger,
	gormDB *gorm.DB,
) ExposedBusinessLogic {
	return &businesslogic.Impl{
		Repository: &infra_repository.Impl{
			GormDB:  gormDB,
			NowFunc: time.Now,
		},
		GenerateChatMessageID: &proc.IDGeneratorImpl{},
		L:                     l,
	}
}
