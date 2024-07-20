package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) GetPhotoStudioUsers(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	q *cgorm.ListQuery,
) ([]*entity.PhotoStudioUser, bool, error) {
	return t.Repository.GetPhotoStudioUsers(ctx, photoStudioID, q)
}

func (t *Impl) GetPhotoStudioUser(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
) (*entity.PhotoStudioUser, error) {
	return t.Repository.GetPhotoStudioUser(ctx, photoStudioID, userID)
}

func (t *Impl) CreatePhotoStudioUser(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	user *common_entity.User,
) (*entity.PhotoStudioUser, error) {
	return t.Repository.CreatePhotoStudioUser(ctx, photoStudioID, user)
}
