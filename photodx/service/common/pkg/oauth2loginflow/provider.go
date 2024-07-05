package oauth2loginflow

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type ProviderID string

const (
	ProviderLINE ProviderID = "line"
)

type Provider struct {
	ClientID     string
	ClientSecret string
}

type Oauth2FetchProfileFunc = func(
	ctx context.Context,
	accessToken string,
) (*entity.User, ResourceOwnerID, error)
