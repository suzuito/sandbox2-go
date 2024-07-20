package usecase

import (
	"context"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type Usecase interface {
	MiddlewareAccessTokenAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOMiddlewareAccessTokenAuthe, error)
	MiddlewarePhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*DTOMiddlewarePhotoStudio, error)
}
