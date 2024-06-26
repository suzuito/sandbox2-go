package inject

import (
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase"
	"github.com/suzuito/sandbox2-go/photodx/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/internal/web/presenter"
	"github.com/suzuito/sandbox2-go/photodx/pkg/environment"
)

func NewWebImpl(
	env *environment.Environment,
	logger *slog.Logger,
	u usecase.Usecase,
) (*web.Impl, error) {
	auth0IssuerURL, err := url.Parse(fmt.Sprintf("https://%s/", env.Auth0Domain))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the issuer url: %w", err)
	}
	auth0Provider := jwks.NewCachingProvider(auth0IssuerURL, 1*time.Minute)
	auth0Validator, err := validator.New(
		auth0Provider.KeyFunc,
		validator.RS256,
		auth0IssuerURL.String(),
		[]string{
			env.Auth0Audience,
		},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &web.Auth0CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set up the jwt validator: %w", err)
	}
	return &web.Impl{
		U:                 u,
		P:                 &presenter.Impl{},
		L:                 logger,
		Auth0Validator:    auth0Validator,
		CorsAllowOrigins:  env.CorsAllowOrigins,
		CorsAllowMethods:  env.CorsAllowMethods,
		CorsAllowHeaders:  env.CorsAllowHeaders,
		CorsExposeHeaders: env.CorsExposeHeaders,
	}, nil
}
