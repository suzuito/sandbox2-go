package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	GetActiveLineLink(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*entity.LineLinkInfo, error)
	ActivateLineLink(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*entity.LineLinkInfo, error)
	DeactivateLineLink(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*entity.LineLinkInfo, error)
	SetLineLinkInfo(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		arg *repository.SetLineLinkInfoArgument,
	) (*entity.LineLinkInfo, error)

	GenerateUserFromLINEProfile(
		ctx context.Context,
		lineLinkInfo *entity.LineLinkInfo,
		lineUserID string,
	) (*common_entity.User, error)
	CreatePhotoStudioUser(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		user *common_entity.User,
	) (*entity.PhotoStudioUser, error)
	GetPhotoStudioUsers(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		q *cgorm.ListQuery,
	) ([]*entity.PhotoStudioUser, bool, error)
	GetPhotoStudioUser(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
	) (*entity.PhotoStudioUser, error)
}
