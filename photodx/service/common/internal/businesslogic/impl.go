package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

type Impl struct {
	Repository     repository.Repository
	SaltRepository repository.SaltRepository
	NowFunc        func() time.Time

	PhotoStudioMemberIDGenerator              proc.RandomStringGenerator
	PhotoStudioMemberInitialPasswordGenerator proc.RandomStringGenerator
	PasswordHasher                            proc.PasswordHasher
	AdminRefreshTokenJWTCreator               auth.JWTCreator
	AdminRefreshTokenJWTVerifier              auth.JWTVerifier
	AdminAccessTokenJWTCreator                auth.JWTCreator
	AdminAccessTokenJWTVerifier               auth.JWTVerifier

	UserRefreshTokenJWTCreator    auth.JWTCreator
	UserRefreshTokenJWTVerifier   auth.JWTVerifier
	UserAccessTokenJWTCreator     auth.JWTCreator
	UserAccessTokenJWTVerifier    auth.JWTVerifier
	OAuth2LoginFlowStateGenerator proc.RandomStringGenerator
	UserIDGenerator               proc.RandomStringGenerator
}
