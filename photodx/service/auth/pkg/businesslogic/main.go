package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/repository"
	infra_saltrepository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/saltrepository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"gorm.io/gorm"
)

func Main(
	gormDB *gorm.DB,
	webPushVAPIDPrivateKey string,
	webPushVAPIDPublicKey string,
) ExposedBusinessLogic {
	return &businesslogic.Impl{
		Repository: &infra_repository.Impl{
			GormDB:  gormDB,
			NowFunc: time.Now,
		},
		SaltRepository:                            &infra_saltrepository.Impl{},
		PasswordHasher:                            &proc.PasswordHasherMD5{},
		PhotoStudioMemberIDGenerator:              &proc.IDGeneratorImpl{},
		PhotoStudioMemberInitialPasswordGenerator: &proc.InitialPasswordGeneratorImpl{},
		WebPushVAPIDPrivateKey:                    webPushVAPIDPrivateKey,
		WebPushVAPIDPublicKey:                     webPushVAPIDPublicKey,
		NowFunc:                                   time.Now,
	}
}
