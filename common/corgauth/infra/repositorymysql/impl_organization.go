package repositorymysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
	"github.com/suzuito/sandbox2-go/common/corgauth/service/repository"
	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) GetOrganization(
	ctx context.Context,
	organizationID entity.OrganizationID,
) (*entity.Organization, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id` FROM `organizations` WHERE `id` = ?",
		organizationID,
	)
	organization := entity.Organization{}
	if err := row.Scan(
		&organization.ID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, terrors.Wrap(&repository.EntityNotFoundError{
				EntityType: repository.EntityTypeOrganization,
				ID:         string(organizationID),
			})
		}
		return nil, terrors.Wrap(err)
	}
	return &organization, nil
}

func (t *Impl) CreateOrganization(
	ctx context.Context,
	organization *entity.Organization,
) (*entity.Organization, error) {
	if _, err := t.GetOrganization(ctx, organization.ID); err == nil {
		return nil, terrors.Wrap(
			&repository.EntityAlreadyExistsError{
				EntityType: repository.EntityTypeOrganization,
				ID:         string(organization.ID),
			},
		)
	}
	if err := csql.WithTransaction(ctx, t.Pool, func(tx csql.TxOrDB) error {
		if _, err := csql.ExecContext(
			ctx,
			tx,
			"INSERT INTO `organizations`(`id`, `active`, `created_at`, `updated_at`) VALUES (?,?,NOW(), NOW())",
			organization.ID,
			organization.Active,
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, terrors.Wrap(err)
	}
	created, err := t.GetOrganization(ctx, organization.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}
