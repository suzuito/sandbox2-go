package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type Repository interface {
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*common_entity.PhotoStudio, error)
	GetPhotoStudios(
		ctx context.Context,
		photoStudioIDs []common_entity.PhotoStudioID,
	) ([]*common_entity.PhotoStudio, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudio *common_entity.PhotoStudio,
	) (*common_entity.PhotoStudio, error)

	GetPhotoStudioMember(
		ctx context.Context,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
	) (*common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, error)
	GetPhotoStudioMembers(
		ctx context.Context,
		photoStudioMemberIDs []common_entity.PhotoStudioMemberID,
	) ([]*common_entity.PhotoStudioMemberWrapper, error)
	ListPhotoStudioMembers(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		listQuery *cgorm.ListQuery,
	) ([]*common_entity.PhotoStudioMemberWrapper, bool, error)
	GetPhotoStudioMemberByEmail(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		email string,
	) (*common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, error)
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		photoStudioMember *common_entity.PhotoStudioMember,
		initialPasswordHashValue string,
		initialRoles []rbac.RoleID,
	) (*common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, error)
	GetPhotoStudioMemberPasswordHashByEmail(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		email string,
	) (string, *common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, error)

	GetLatestPhotoStudioMemberWebPushSubscriptions(
		ctx context.Context,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
	) ([]*entity.PhotoStudioMemberWebPushSubscription, error)
	UpdateOrCreateUserWebPushSubscription(
		ctx context.Context,
		s *entity.PhotoStudioMemberWebPushSubscription,
	) (*entity.PhotoStudioMemberWebPushSubscription, error)
}
