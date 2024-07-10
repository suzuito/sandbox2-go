package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOGetCallback struct {
	User         *entity.User
	RefreshToken string
	AccessToken  string
}

func (t *Impl) GetCallback(
	ctx context.Context,
	code string,
	stateCode oauth2loginflow.StateCode,
) (*DTOGetCallback, error) {
	state, err := t.BusinessLogic.CallbackVerifyState(ctx, stateCode)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	var tokenRequest *http.Request
	var fetchProfile oauth2loginflow.Oauth2FetchProfileFunc
	switch state.ProviderID {
	case oauth2loginflow.ProviderLINE:
		body := url.Values{}
		body.Set("code", code)
		body.Set("grant_type", "authorization_code")
		body.Set("redirect_uri", state.RedirectURL)
		body.Set("client_id", t.OAuth2ProviderLINE.ClientID)
		body.Set("client_secret", t.OAuth2ProviderLINE.ClientSecret)
		tokenRequest, _ = http.NewRequest(http.MethodPost, "https://api.line.me/oauth2/v2.1/token", strings.NewReader(body.Encode()))
		tokenRequest.Header.Set("content-type", "application/x-www-form-urlencoded")
		fetchProfile = fetchLINEProfile
	default:
		return nil, terrors.Wrapf("unsupported provider %s", state.ProviderID)
	}
	accessTokenFromProvider, err := t.BusinessLogic.FetchAccessToken(ctx, tokenRequest)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	user, err := t.BusinessLogic.FetchProfileAndCreateUserIfNotExists(ctx, accessTokenFromProvider, state.ProviderID, fetchProfile)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.BusinessLogic.CreateUserRefreshToken(ctx, user.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	accessToken, err := t.BusinessLogic.CreateUserAccessToken(ctx, user.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetCallback{
		User:         user,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func fetchLINEProfile(
	ctx context.Context,
	accessToken string,
) (*entity.User, oauth2loginflow.ResourceOwnerID, error) {
	cli := http.DefaultClient
	// fetch profile
	urlProfile, _ := url.Parse("https://api.line.me/v2/profile")
	req, _ := http.NewRequest(
		http.MethodGet,
		urlProfile.String(),
		nil,
	)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resProfile, err := cli.Do(req)
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	defer resProfile.Body.Close()
	if resProfile.StatusCode != http.StatusOK {
		bodyBytesProfile, _ := io.ReadAll(resProfile.Body)
		return nil, "", terrors.Wrapf("http error %d %s", resProfile.StatusCode, string(bodyBytesProfile)) // TODO
	}
	bodyBytesProfile, err := io.ReadAll(resProfile.Body)
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	profile := struct {
		UserID      string `json:"userId"`
		DisplayName string `json:"displayName"`
		PictureURL  string `json:"pictureUrl"`
	}{}
	if err := json.Unmarshal(bodyBytesProfile, &profile); err != nil {
		return nil, "", terrors.Wrap(err)
	}
	return &entity.User{
		Name:            profile.DisplayName,
		ProfileImageURL: profile.PictureURL,
	}, oauth2loginflow.ResourceOwnerID(profile.UserID), nil
}
