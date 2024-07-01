package service

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/auth"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/proc"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

type Impl struct {
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
