package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

func (t *Impl) GetPhotoStudioMemberRoles(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) ([]*rbac.Role, error) {
	return t.Repository.GetPhotoStudioMemberRoles(ctx, photoStudioMemberID)
}
