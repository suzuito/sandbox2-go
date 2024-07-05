package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type BusinessLogic interface {
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

	// impl_admin_access_token.go
	CreateAdminAccessToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyAdminAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.AdminPrincipal, error)

	// impl_admin_refresh_token.go
	CreateAdminRefreshToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyAdminRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.AdminPrincipalRefreshToken, error)
}
