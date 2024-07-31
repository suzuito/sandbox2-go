package businesslogic

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/gateway/mail"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
)

type Impl struct {
	Repository                                repository.Repository
	NowFunc                                   func() time.Time
	UserRefreshTokenJWTCreator                auth.JWTCreator
	UserRefreshTokenJWTVerifier               auth.JWTVerifier
	UserAccessTokenJWTCreator                 auth.JWTCreator
	UserAccessTokenJWTVerifier                auth.JWTVerifier
	OAuth2LoginFlowStateGenerator             proc.RandomStringGenerator
	PromoteGuestUserConfirmationCodeGenerator proc.RandomStringGenerator
	UserIDGenerator                           proc.RandomStringGenerator
	UserMailSender                            mail.UserMailSender
	PasswordSalt                              string
	PasswordHasher                            proc.PasswordHasher
	WebPushVAPIDPublicKey                     string
	WebPushVAPIDPrivateKey                    string
}
