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
	refreshTokenJWTCreator auth.JWTCreator,
	refreshTokenJWTVerifier auth.JWTVerifier,
	accessTokenJWTCreator auth.JWTCreator,
	accessTokenJWTVerifier auth.JWTVerifier,
	nowFunc func() time.Time,
) businesslogic.BusinessLogic {
	return &businesslogic_internal.BusinessLogicImpl{
		PhotoStudioMemberIDGenerator:              photoStudioMemberIDGenerator,
		PhotoStudioMemberInitialPasswordGenerator: photoStudioMemberInitialPasswordGenerator,
		PasswordHasher:                            passwordHasher,
		Repository:                                repository,
		SaltRepository:                            saltRepository,
		RefreshTokenJWTCreator:                    refreshTokenJWTCreator,
		RefreshTokenJWTVerifier:                   refreshTokenJWTVerifier,
		AccessTokenJWTCreator:                     accessTokenJWTCreator,
		AccessTokenJWTVerifier:                    accessTokenJWTVerifier,
		NowFunc:                                   nowFunc,
	}
}
