package inject

import (
	"time"

	businesslogic_internal "github.com/suzuito/sandbox2-go/photodx/service/common/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func NewBusinessLogic(
	photoStudioMemberIDGenerator proc.RandomStringGenerator,
	photoStudioMemberInitialPasswordGenerator proc.RandomStringGenerator,
	passwordHasher proc.PasswordHasher,
	repository repository.Repository,
	saltRepository repository.SaltRepository,
	adminRefreshTokenJWTCreator auth.JWTCreator,
	adminRefreshTokenJWTVerifier auth.JWTVerifier,
	adminAccessTokenJWTCreator auth.JWTCreator,
	adminAccessTokenJWTVerifier auth.JWTVerifier,
	userAccessTokenJWTVerifier auth.JWTVerifier,
	nowFunc func() time.Time,
) businesslogic.BusinessLogic {
	return &businesslogic_internal.Impl{
		Repository:     repository,
		SaltRepository: saltRepository,
		NowFunc:        nowFunc,

		PhotoStudioMemberIDGenerator:              photoStudioMemberIDGenerator,
		PhotoStudioMemberInitialPasswordGenerator: photoStudioMemberInitialPasswordGenerator,
		PasswordHasher:                            passwordHasher,
		AdminRefreshTokenJWTCreator:               adminRefreshTokenJWTCreator,
		AdminRefreshTokenJWTVerifier:              adminRefreshTokenJWTVerifier,
		AdminAccessTokenJWTCreator:                adminAccessTokenJWTCreator,
		AdminAccessTokenJWTVerifier:               adminAccessTokenJWTVerifier,

		// TODO ↓出鱈目。後でちゃんとする。
		UserAccessTokenJWTVerifier:    adminAccessTokenJWTVerifier,
		UserIDGenerator:               photoStudioMemberIDGenerator,
		UserAccessTokenJWTCreator:     adminAccessTokenJWTCreator,
		UserRefreshTokenJWTCreator:    adminRefreshTokenJWTCreator,
		UserRefreshTokenJWTVerifier:   adminRefreshTokenJWTVerifier,
		OAuth2LoginFlowStateGenerator: photoStudioMemberIDGenerator,
	}
}
