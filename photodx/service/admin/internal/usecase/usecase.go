package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type Usecase interface {
	MiddlewareAccessTokenAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOMiddlewareAccessTokenAuthe, error)
	APIMiddlewarePhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
	) (*DTOAPIMiddlewarePhotoStudio, error)
	APIGetInit(
		ctx context.Context,
		principal entity.AdminPrincipalAccessToken,
	) (*DTOAPIGetInit, error)

	APIGetLINELink(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
	) (*DTOAPIGetLINELink, error)
	APIPostLINELink(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
	) (*DTOAPIPostLINELink, error)
	APIDeleteLINELink(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
	) error
	APIPutLINELinkMessagingAPIChannelSecret(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
		secret string,
	) (*DTOAPIPutLINELinkMessagingAPIChannelSecret, error)
}
