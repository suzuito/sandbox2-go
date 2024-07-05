package usecase

import (
	"context"
	"net/url"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
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
