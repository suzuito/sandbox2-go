package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

type Service interface {
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
	) (*entity.PhotoStudio, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		name string,
	) (*entity.PhotoStudio, error)

	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		name string,
	) (*entity.PhotoStudioMember, string, error)
	GetPhotoStudioMember(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (*entity.PhotoStudioMember, error)
	GetPhotoStudioMemberByEmail(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
	) (*entity.PhotoStudioMember, error)
	SendPhotoStudioMemberInvitation(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) error
	VerifyPhotoStudioMemberPassword(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		password string,
	) error

	GetPhotoStudioMemberRoles(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) ([]*rbac.Role, error)

	CreateAccessToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.Principal, error)

	CreateRefreshToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.PrincipalRefreshToken, error)
}
