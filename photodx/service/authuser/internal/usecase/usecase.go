package usecase

import (
	"context"
	"net/url"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
)

type Usecase interface {
	GetAuthorizeURLLINE(
		ctx context.Context,
		callbackURL *url.URL,
		oauth2RedirectURL *url.URL,
	) (*DTOGetAuthorizeLINE, error)
	GetCallback(
		ctx context.Context,
		code string,
		stateCode oauth2loginflow.StateCode,
	) (*DTOGetCallback, error)
}
