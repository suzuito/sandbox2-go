package auth

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type JWTVerifierAuth0 struct {
	Auth0Validator *validator.Validator
}

var signingMethodAuth0 = jwt.SigningMethodRS256

func (t *JWTVerifierAuth0) VerifyJWTToken(
	ctx context.Context,
	tokenString string,
	zeroClaims jwt.Claims,
) (jwt.Claims, error) {
	claims, err := t.Auth0Validator.ValidateToken(ctx, tokenString)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	validatedClaims, ok := claims.(*validator.ValidatedClaims)
	if !ok {
		return nil, terrors.Wrapf("cannot convert claims to *validator.ValdatedClaims")
	}
	customClaims, ok := validatedClaims.CustomClaims.(*auth0CustomClaims)
	if !ok {
		return nil, terrors.Wrapf("cannot convert custom claims to *auth0CustomClaims")
	}
	roles := []*rbac.Role{}
	for _, scope := range strings.Fields(customClaims.Scope) {
		role, exists := rbac.AvailablePredefinedRoles[rbac.RoleID(scope)]
		if !exists {
			continue
		}
		roles = append(roles, role)
	}
	jwtClaims := JWTClaimsUserAccessToken{
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: validatedClaims.RegisteredClaims.Subject,
		},
	}
	return &jwtClaims, nil
}

type auth0CustomClaims struct {
	Scope string `json:"scope"`
}

func (t *auth0CustomClaims) Validate(context.Context) error {
	return nil
}

func NewJWTVerifierAuth0(
	auth0Domain string,
	auth0Audience string,
) (*JWTVerifierAuth0, error) {
	auth0IssuerURL, err := url.Parse(fmt.Sprintf("https://%s/", auth0Domain))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the issuer url: %w", err)
	}
	auth0Provider := jwks.NewCachingProvider(auth0IssuerURL, 1*time.Minute)
	auth0Validator, err := validator.New(
		auth0Provider.KeyFunc,
		validator.RS256,
		auth0IssuerURL.String(),
		[]string{
			auth0Audience,
		},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &auth0CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set up the jwt validator: %w", err)
	}
	return &JWTVerifierAuth0{
		Auth0Validator: auth0Validator,
	}, nil
}
