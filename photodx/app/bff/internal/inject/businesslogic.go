package inject

import (
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	common_inject "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/inject"
)

func NewBusinessLogic(
	env *environment.Environment,
	logger *slog.Logger,
) (businesslogic.BusinessLogic, error) {
	adminAccessTokenJWTPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(env.JWTAdminAccessTokenSigningPublicKey))
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	b := common_inject.NewBusinessLogic(
		&auth.JWTVerifiers{
			Verifiers: []auth.JWTVerifier{
				&auth.JWTVerifierRS{
					PublicKey: adminAccessTokenJWTPublicKey,
				},
			},
		},
	)
	return b, nil
}
