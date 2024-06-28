package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) VerifyPhotoStudioMemberPassword(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
	password string,
) error {
	salt, err := t.SaltRepository.Get(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	hashInput := generatePasswordHash(salt, password)
	hashInDB, err := t.Repository.GetPhotoStudioMemberPasswordHashByEmail(ctx, photoStudioID, email)
	if err != nil {
		return terrors.Wrap(err)
	}
	if hashInput != hashInDB {
		return ErrPasswordMismatch
	}
	return nil
}
