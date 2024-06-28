package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

type Usecase interface {
	APIMiddlewareAuthAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOAPIMiddlewareAuthAuthe, error)
	APIMiddlewarePhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
	) (*DTOAPIMiddlewarePhotoStudio, error)
	AuthPostLogin(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		password string,
	) (*DTOAuthPostLogin, error)
	APIPostPhotoStudioMembers(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		name string,
	) (*DTOAPIPostPhotoStudioMembers, error)
	APIPostPhotoStudios(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		name string,
	) (*DTOAPIPostPhotoStudios, error)
	SuperPostInit(ctx context.Context) (*DTOSuperPostInit, error)
}
