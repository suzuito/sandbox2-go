package usecase

import (
	"context"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
)

type DTOGetAuthorizeLINE struct {
	AuthorizeURL *url.URL
}

// var clientID = "2005761043"
// var clientSecret = "3250327d6ab0c0f92938d37e6ff87750"

func (t *Impl) GetAuthorizeURLLINE(
	ctx context.Context,
	callbackURL *url.URL,
	oauth2RedirectURL *url.URL,
) (*DTOGetAuthorizeLINE, error) {
	state, err := t.BusinessLogic.CreateOAuth2State(ctx, oauth2loginflow.ProviderLINE, callbackURL, oauth2RedirectURL)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	authURL, _ := url.Parse("https://access.line.me/oauth2/v2.1/authorize")
	authURLQuery := authURL.Query()
	authURLQuery.Set("response_type", "code")
	authURLQuery.Set("client_id", t.OAuth2ProviderLINE.ClientID)
	authURLQuery.Set("redirect_uri", oauth2RedirectURL.String())
	authURLQuery.Set("state", string(state.Code))
	authURLQuery.Set("scope", "profile")
	authURLQuery.Set("ui_locales", "ja-JP")
	authURLQuery.Set("initial_amr_display", "lineqr")
	authURL.RawQuery = authURLQuery.Encode()
	return &DTOGetAuthorizeLINE{
		AuthorizeURL: authURL,
	}, nil
}
