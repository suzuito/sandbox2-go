package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

type BusinessLogicImpl struct {
	PhotoStudioMemberIDGenerator              proc.RandomStringGenerator
	PhotoStudioMemberInitialPasswordGenerator proc.RandomStringGenerator
	PasswordHasher                            proc.PasswordHasher
	Repository                                repository.Repository
	SaltRepository                            repository.SaltRepository
	RefreshTokenJWTCreator                    auth.JWTCreator
	RefreshTokenJWTVerifier                   auth.JWTVerifier
	AccessTokenJWTCreator                     auth.JWTCreator
	AccessTokenJWTVerifier                    auth.JWTVerifier
	NowFunc                                   func() time.Time
}
