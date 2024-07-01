package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

type Repository interface {
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
	) (*entity.PhotoStudio, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudio *entity.PhotoStudio,
	) (*entity.PhotoStudio, error)

	GetPhotoStudioMember(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	GetPhotoStudioMemberByEmail(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		photoStudioMember *entity.PhotoStudioMember,
		initialPasswordHashValue string,
		initialRoles []rbac.RoleID,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	GetPhotoStudioMemberPasswordHashByEmail(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
	) (string, *entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
}
