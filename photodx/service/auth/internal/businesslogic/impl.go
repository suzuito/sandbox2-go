package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
)

type Impl struct {
	Repository                                repository.Repository
	PasswordSalt                              string
	PasswordHasher                            proc.PasswordHasher
	PhotoStudioMemberIDGenerator              proc.RandomStringGenerator
	PhotoStudioMemberInitialPasswordGenerator proc.RandomStringGenerator
	AdminRefreshTokenJWTCreator               auth.JWTCreator
	AdminRefreshTokenJWTVerifier              auth.JWTVerifier
	AdminAccessTokenJWTCreator                auth.JWTCreator
	AdminAccessTokenJWTVerifier               auth.JWTVerifier
	WebPushVAPIDPublicKey                     string
	WebPushVAPIDPrivateKey                    string
	NowFunc                                   func() time.Time
}
