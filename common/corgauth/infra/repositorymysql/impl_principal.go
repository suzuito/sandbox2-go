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

func (t *Impl) GetPrincipal(
	ctx context.Context,
	principalID entity.PrincipalID,
) (*entity.Principal, []*entity.Role, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id`, `email`, `active` FROM `principals` WHERE `id` = ?",
		principalID,
	)
	principal := entity.Principal{}
	if err := row.Scan(
		&principal.ID,
		&principal.Email,
		&principal.Active,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, terrors.Wrap(&repository.EntityNotFoundError{
				EntityType: repository.EntityTypePrincipal,
				ID:         string(principal.ID),
			})
		}
		return nil, nil, terrors.Wrap(err)
	}
	roles, err := getPrincipalRoles(
		ctx,
		t.Pool,
		[]entity.PrincipalID{principal.ID},
	)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	returnedRoles, exists := roles[principal.ID]
	if !exists {
		returnedRoles = []*entity.Role{}
	}
	return &principal, returnedRoles, nil
}

func (t *Impl) GetPrincipalByEmail(
	ctx context.Context,
	organizationID entity.OrganizationID,
	email string,
) (*entity.Principal, []*entity.Role, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id` FROM `principals` WHERE `organization_id` = ? AND `email` = ?",
		organizationID,
		email,
	)
	principalID := entity.PrincipalID("")
	if err := row.Scan(
		&principalID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, terrors.Wrap(&repository.EntityNotFoundError{
				EntityType: repository.EntityTypePrincipal,
				ID:         "email",
			})
		}
		return nil, nil, terrors.Wrap(err)
	}
	return t.GetPrincipal(ctx, principalID)
}

func (t *Impl) CreatePrincipal(
	ctx context.Context,
	organizationID entity.OrganizationID,
	principal *entity.Principal,
	initialPasswordHash string,
	initialRoles []entity.Role,
) (*entity.Principal, []*entity.Role, error) {
	if _, _, err := t.GetPrincipal(ctx, principal.ID); err == nil {
		return nil, nil, terrors.Wrap(
			&repository.EntityAlreadyExistsError{
				EntityType: repository.EntityTypePrincipal,
				ID:         string(principal.ID),
			},
		)
	}
	if err := csql.WithTransaction(ctx, t.Pool, func(tx csql.TxOrDB) error {
		if _, err := csql.ExecContext(
			ctx,
			tx,
			"INSERT INTO `principals`(`id`, `organization_id`, `email`, `active`, `created_at`, `updated_at`) VALUES(?, ?, ?, ?, NOW(), NOW())",
			principal.ID,
			organizationID,
			principal.Email,
			principal.Active,
		); err != nil {
			return terrors.Wrap(err)
		}
		if _, err := csql.ExecContext(
			ctx,
			tx,
			"INSERT INTO `principal_password_hashes`(`principal_id`, `value`, `created_at`, `updated_at`) VALUES(?, ?, NOW(), NOW())",
			principal.ID,
			initialPasswordHash,
		); err != nil {
			return terrors.Wrap(err)
		}
		if _, err := setPrincipalRoles(
			ctx,
			tx,
			principal.ID,
			initialRoles,
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	return t.GetPrincipal(ctx, principal.ID)
}

func (t *Impl) GetPrincipalPasswordHash(
	ctx context.Context,
	principalID entity.PrincipalID,
) (string, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `value` FROM `principal_password_hashes` WHERE `principal_id` = ?",
		principalID,
	)
	hashValue := ""
	if err := row.Scan(&hashValue); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", terrors.Wrap(&repository.EntityNotFoundError{
				EntityType: "PhotoStudioMemberHashValues",
			})
		}
		return "", terrors.Wrap(err)
	}
	return hashValue, nil
}
