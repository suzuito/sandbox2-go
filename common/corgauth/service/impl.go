package service

import (
	"time"

	"github.com/suzuito/sandbox2-go/common/corgauth/service/jwttoken"
	"github.com/suzuito/sandbox2-go/common/corgauth/service/proc"
	"github.com/suzuito/sandbox2-go/common/corgauth/service/repository"
)

type Impl struct {
	Repository              repository.Repository
	SaltRepository          repository.SaltRepository
	GeneratePrincipalID     proc.GenerateIDFunc
	GeneratePassword        proc.GeneratePasswordFunc
	GeneratePasswordHash    proc.GeneratePasswordHashFunc
	RefreshTokenJWTCreator  jwttoken.JWTCreator
	RefreshTokenJWTVerifier jwttoken.JWTVerifier
	AccessTokenJWTCreator   jwttoken.JWTCreator
	AccessTokenJWTVerifier  jwttoken.JWTVerifier
	NowFunc                 func() time.Time
}
