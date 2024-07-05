package businesslogic

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func (t *Impl) CreateOAuth2State(
	ctx context.Context,
	providerID oauth2loginflow.ProviderID,
	callbackURL *url.URL,
	oauth2RedirectURL *url.URL,
) (*oauth2loginflow.State, error) {
	stateCode, err := t.OAuth2LoginFlowStateGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	state := oauth2loginflow.State{
		Code:        oauth2loginflow.StateCode(stateCode),
		ProviderID:  oauth2loginflow.ProviderLINE,
		RedirectURL: oauth2RedirectURL.String(),
		CallbackURL: callbackURL.String(),
		ExpiresAt:   t.NowFunc().Add(time.Second * 600),
	}
	created, err := t.Repository.CreateOAuth2State(ctx, &state)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}

func (t *Impl) CallbackVerifyState(
	ctx context.Context,
	stateCode oauth2loginflow.StateCode,
) (*oauth2loginflow.State, error) {
	state, err := t.Repository.GetAndDeleteOAuth2State(ctx, stateCode)
	if err != nil {
		var noEntryError *repository.NoEntryError
		if errors.As(err, &noEntryError) {
			return nil, terrors.Wrap(oauth2loginflow.ErrStateMismatch)
		}
		return nil, terrors.Wrap(err)
	}
	now := t.NowFunc()
	if now.After(state.ExpiresAt) {
		return nil, terrors.Wrap(oauth2loginflow.ErrStateMismatch)
	}
	return state, nil
}

func (t *Impl) FetchAccessToken(
	ctx context.Context,
	req *http.Request,
) (string, error) {
	cli := http.DefaultClient
	res, err := cli.Do(req)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		bodyBytesProfile, _ := io.ReadAll(res.Body)
		return "", terrors.Wrapf("http error %d %s", res.StatusCode, string(bodyBytesProfile)) // TODO
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	oauth2Res := struct {
		RefreshToken string        `json:"refresh_token"`
		AccessToken  string        `json:"access_token"`
		Scope        string        `json:"scope"`
		TokenType    string        `json:"token_type"`
		ExpiresIn    time.Duration `json:"expires_in"`
		IDToken      string        `json:"id_token"`
	}{}
	if err := json.Unmarshal(bodyBytes, &oauth2Res); err != nil {
		return "", terrors.Wrap(err)
	}
	return oauth2Res.AccessToken, nil
}

func (t *Impl) FetchProfileAndCreateUserIfNotExists(
	ctx context.Context,
	accessToken string,
	providerID oauth2loginflow.ProviderID,
	fetchProfile oauth2loginflow.Oauth2FetchProfileFunc,
) (*entity.User, error) {
	user, resourceOwnerID, err := fetchProfile(ctx, accessToken)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	existingUser, err := t.Repository.GetUserByResourceOwnerID(ctx, providerID, resourceOwnerID)
	if err == nil {
		return existingUser, nil
	}
	var noEntryError *repository.NoEntryError
	if !errors.As(err, &noEntryError) {
		return nil, terrors.Wrap(err)
	}
	userID, err := t.UserIDGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	user.ID = entity.UserID(userID)
	createdUser, err := t.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return createdUser, nil
}
