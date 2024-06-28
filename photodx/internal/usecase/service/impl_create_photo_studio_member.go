package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) CreatePhotoStudioMember(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
	name string,
) (*entity.PhotoStudioMember, string, error) {
	id, err := t.PhotoStudioMemberIDGenerator.Gen()
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	member := entity.PhotoStudioMember{
		ID:     entity.PhotoStudioMemberID(id),
		Email:  email,
		Name:   name,
		Active: false,
	}
	if err := member.Validate(); err != nil {
		return nil, "", terrors.Wrap(err)
	}
	if _, err := t.Repository.GetPhotoStudioMemberByEmail(ctx, photoStudioID, member.Email); err == nil {
		return nil, "", terrors.Wrap(&repository.DuplicateEntryError{EntryType: "PhotoStudioMember"})
	}
	initialPassword, err := t.PhotoStudioMemberInitialPasswordGenerator.Gen()
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	salt, err := t.SaltRepository.Get(ctx)
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	initialPasswordHashValue := generatePasswordHash(salt, initialPassword)
	created, err := t.Repository.CreatePhotoStudioMember(
		ctx,
		photoStudioID,
		&member,
		string(initialPasswordHashValue),
	)
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	return created, initialPassword, nil
}
