package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) CreateOrganization(
	ctx context.Context,
	organizationID entity.OrganizationID,
) (*entity.Organization, error) {
	organization := entity.Organization{
		ID:     organizationID,
		Active: false,
	}
	if err := organization.Validate(); err != nil {
		return nil, terrors.Wrap(err)
	}
	created, err := t.Repository.CreateOrganization(ctx, &organization)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}
