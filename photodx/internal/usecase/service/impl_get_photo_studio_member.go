package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) GetPhotoStudioMemberByEmail(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
) (*entity.PhotoStudioMember, error) {
	member, err := t.Repository.GetPhotoStudioMemberByEmail(ctx, photoStudioID, email)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return member, nil
}
