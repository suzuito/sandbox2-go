package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

type Service interface {
	// impl_photo_studio.go
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
	) (*entity.PhotoStudio, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		name string,
	) (*entity.PhotoStudio, error)

	// impl_photo_studio_member.go
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		name string,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, string, error)
	GetPhotoStudioMember(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	VerifyPhotoStudioMemberPassword(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		password string,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)

	// impl_access_token.go
	CreateAccessToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.Principal, error)

	// impl_refresh_token.go
	CreateRefreshToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.PrincipalRefreshToken, error)
}
