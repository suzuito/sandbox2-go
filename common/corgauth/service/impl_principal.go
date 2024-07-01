package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) CreatePrincipal(
	ctx context.Context,
	organizationID entity.OrganizationID,
	email string,
	initialRoles []entity.Role,
) (*entity.Principal, string, error) {
	_, err := t.Repository.GetOrganization(ctx, organizationID)
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	id, err := t.GeneratePrincipalID()
	if err != nil {
		return nil, "", terrors.Wrapf("GeneratePrincipalID is failed : %w", err)
	}
	salt, err := t.SaltRepository.Get(ctx)
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	password, err := t.GeneratePassword()
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	passwordHash := t.GeneratePasswordHash(password, salt)
	principalID := entity.PrincipalID(id)
	principal := entity.Principal{
		ID:     principalID,
		Email:  email,
		Active: false,
	}
	principalCreated, _, err := t.Repository.CreatePrincipal(
		ctx,
		organizationID,
		&principal,
		passwordHash,
		initialRoles,
	)
	if err != nil {
		return nil, "", terrors.Wrap(err)
	}
	return principalCreated, password, nil
}

func (t *Impl) VerifyPrincipalPassword(
	ctx context.Context,
	organizationID entity.OrganizationID,
	email string,
	password string,
) (*entity.Principal, []*entity.Role, error) {
	salt, err := t.SaltRepository.Get(ctx)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	principal, roles, err := t.Repository.GetPrincipalByEmail(ctx, organizationID, email)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	passwordHashInDB, err := t.Repository.GetPrincipalPasswordHash(ctx, principal.ID)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	passwordHashInput := t.GeneratePasswordHash(password, salt)
	if passwordHashInDB != passwordHashInput {
		return nil, nil, terrors.Wrap(ErrPasswordMismatch)
	}
	return principal, roles, nil
}
