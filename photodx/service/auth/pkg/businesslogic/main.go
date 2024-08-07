package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
	infra_repository "github.com/suzuito/sandbox2-go/photodx/service/auth/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"gorm.io/gorm"
)

func Main(
	gormDB *gorm.DB,
	passwordSalt string,
	webPushVAPIDPrivateKey string,
	webPushVAPIDPublicKey string,
) ExposedBusinessLogic {
	return &businesslogic.Impl{
		Repository: &infra_repository.Impl{
			GormDB:  gormDB,
			NowFunc: time.Now,
		},
		PasswordSalt:                              passwordSalt,
		PasswordHasher:                            &proc.PasswordHasherMD5{},
		PhotoStudioMemberIDGenerator:              &proc.IDGeneratorImpl{},
		PhotoStudioMemberInitialPasswordGenerator: &proc.InitialPasswordGeneratorImpl{},
		WebPushVAPIDPrivateKey:                    webPushVAPIDPrivateKey,
		WebPushVAPIDPublicKey:                     webPushVAPIDPublicKey,
		NowFunc:                                   time.Now,
	}
}
