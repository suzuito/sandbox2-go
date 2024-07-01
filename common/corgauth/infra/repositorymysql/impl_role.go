package repositorymysql

import (
	"context"
	"strings"

	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func getPrincipalRoles(
	ctx context.Context,
	txOrDB csql.TxOrDB,
	principalIDs []entity.PrincipalID,
) (map[entity.PrincipalID][]*entity.Role, error) {
	csql.ToAnySlice(principalIDs)
	rows, err := csql.QueryContext(
		ctx,
		txOrDB,
		"SELECT `principal_id`, `role_id` FROM `principal_roles` WHERE "+csql.SqlIn("principal_id", principalIDs),
		csql.ToAnySlice(principalIDs)...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rows.Close()
	m := map[entity.PrincipalID][]*entity.Role{}
	for rows.Next() {
		role := entity.Role{}
		principalID := entity.PrincipalID("")
		if err := rows.Scan(
			&principalID,
			&role.ID,
		); err != nil {
			return nil, terrors.Wrap(err)
		}
		roles, exists := m[principalID]
		if !exists {
			roles = []*entity.Role{}
		}
		roles = append(roles, &role)
		m[principalID] = roles
	}
	return m, nil
}

func setPrincipalRoles(
	ctx context.Context,
	txOrDB csql.TxOrDB,
	principalID entity.PrincipalID,
	roles []entity.Role,
) ([]*entity.Role, error) {
	valueStrings := []string{}
	args := []any{}
	for _, roleID := range roles {
		valueStrings = append(valueStrings, "(?,?,NOW())")
		args = append(args, principalID, roleID)
	}
	_, err := csql.ExecContext(
		ctx,
		txOrDB,
		"INSERT IGNORE INTO `principal_roles`(`principal_id`, `role_id`, `created_at`) VALUES "+
			strings.Join(valueStrings, ","),
		args...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	row := csql.QueryRowContext(
		ctx,
		txOrDB,
		"SELECT COUNT(*) FROM `principal_roles` WHERE `principal_id` = ?",
		principalID,
	)
	count := -1
	if err := row.Scan(&count); err != nil {
		return nil, terrors.Wrap(err)
	}
	if count > entity.MaxRolesPerPrincipal {
		return nil, terrors.Wrapf("number of roles is over max : len=%d max=%d", count, entity.MaxRolesPerPrincipal)
	}
	roleMap, err := getPrincipalRoles(
		ctx,
		txOrDB,
		[]entity.PrincipalID{
			principalID,
		},
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	rolesReturned, exists := roleMap[principalID]
	if !exists {
		rolesReturned = []*entity.Role{}
	}
	return rolesReturned, nil
}
