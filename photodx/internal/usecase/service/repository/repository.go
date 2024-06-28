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

	GetPhotoStudioMemberByEmail(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
	) (*entity.PhotoStudioMember, error)
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		photoStudioMember *entity.PhotoStudioMember,
		initialPasswordHashValue string,
	) (*entity.PhotoStudioMember, error)
	GetPhotoStudioMemberPasswordHashByEmail(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
	) (string, error)
	GetPhotoStudioMemberRoles(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) ([]*rbac.Role, error)
}
