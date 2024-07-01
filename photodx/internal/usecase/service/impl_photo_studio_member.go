package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

func (t *Impl) GetPhotoStudioMemberByEmail(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	member, roles, photoStudio, err := t.Repository.GetPhotoStudioMemberByEmail(ctx, photoStudioID, email)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	return member, roles, photoStudio, nil
}

func (t *Impl) GetPhotoStudioMember(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	member, roles, photoStudio, err := t.Repository.GetPhotoStudioMember(ctx, photoStudioMemberID)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	return member, roles, photoStudio, nil
}

func (t *Impl) CreatePhotoStudioMember(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
	name string,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, string, error) {
	id, err := t.PhotoStudioMemberIDGenerator.Gen()
	if err != nil {
		return nil, nil, nil, "", terrors.Wrap(err)
	}
	member := entity.PhotoStudioMember{
		ID:     entity.PhotoStudioMemberID(id),
		Email:  email,
		Name:   name,
		Active: false,
	}
	if err := member.Validate(); err != nil {
		return nil, nil, nil, "", terrors.Wrap(err)
	}
	initialPassword, err := t.PhotoStudioMemberInitialPasswordGenerator.Gen()
	if err != nil {
		return nil, nil, nil, "", terrors.Wrap(err)
	}
	salt, err := t.SaltRepository.Get(ctx)
	if err != nil {
		return nil, nil, nil, "", terrors.Wrap(err)
	}
	initialPasswordHashValue := t.PasswordHasher.Gen(salt, initialPassword)
	created, roles, photoStudio, err := t.Repository.CreatePhotoStudioMember(
		ctx,
		photoStudioID,
		&member,
		string(initialPasswordHashValue),
		[]rbac.RoleID{
			rbac.RoleSuperUser.ID,
		},
	)
	if err != nil {
		return nil, nil, nil, "", terrors.Wrap(err)
	}
	return created, roles, photoStudio, initialPassword, nil
}

func (t *Impl) VerifyPhotoStudioMemberPassword(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
	password string,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	salt, err := t.SaltRepository.Get(ctx)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	hashInput := t.PasswordHasher.Gen(salt, password)
	hashInDB, member, roles, photoStudio, err := t.Repository.GetPhotoStudioMemberPasswordHashByEmail(ctx, photoStudioID, email)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	if hashInput != hashInDB {
		return nil, nil, nil, ErrPasswordMismatch
	}
	return member, roles, photoStudio, nil
}
