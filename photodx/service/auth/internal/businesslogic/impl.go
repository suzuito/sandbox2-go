package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

type Impl struct {
	Repository                                repository.Repository
	SaltRepository                            repository.SaltRepository
	PasswordHasher                            proc.PasswordHasher
	PhotoStudioMemberIDGenerator              proc.RandomStringGenerator
	PhotoStudioMemberInitialPasswordGenerator proc.RandomStringGenerator
	AdminRefreshTokenJWTCreator               auth.JWTCreator
	AdminRefreshTokenJWTVerifier              auth.JWTVerifier
	AdminAccessTokenJWTCreator                auth.JWTCreator
	AdminAccessTokenJWTVerifier               auth.JWTVerifier
	NowFunc                                   func() time.Time
}
