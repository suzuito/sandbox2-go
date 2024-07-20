package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type Usecase interface {
	MiddlewareAccessTokenAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOMiddlewareAccessTokenAuthe, error)

	APIGetLINELink(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
	) (*DTOAPIGetLINELink, error)
	APIPutLINELinkActivate(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
	) (*DTOAPIPutLINELinkActivate, error)
	APIPutLINELinkDeactivate(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
	) (*DTOAPIPutLINELinkDeactivate, error)
	APIPutLINELink(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
		arg *repository.SetLineLinkInfoArgument,
	) (*DTOAPIPutLINELinkMessagingAPIChannelSecret, error)

	APIPostLineMessagingAPIWebhook(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		body []byte,
		xLINESignature string,
		skipVerifySignagure bool,
	) error
}
