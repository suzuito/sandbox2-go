package repository

import (
	"context"
	"strings"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/auth"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/auth/predefined"
)

func (t *Impl) GetPhotoStudioMemberRoles(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) ([]*rbac.Role, error) {
	return getPhotoStudioMemberRoles(ctx, t.Pool, photoStudioMemberID)
}

func getPhotoStudioMemberRoles(
	ctx context.Context,
	txOrDB csql.TxOrDB,
	photoStudioMemberID entity.PhotoStudioMemberID,
) ([]*rbac.Role, error) {
	rows, err := csql.QueryContext(
		ctx,
		txOrDB,
		"SELECT `role_id` FROM `photo_studio_member_roles` WHERE `photo_studio_member_id` = ?",
		photoStudioMemberID,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rows.Close()
	roles := []*rbac.Role{}
	for rows.Next() {
		role := rbac.Role{}
		if err := rows.Scan(
			&role.ID,
		); err != nil {
			return nil, terrors.Wrap(err)
		}
		predefinedRole, exists := predefined.AvailablePredefinedRoles[role.ID]
		if !exists {
			continue
		}
		roles = append(roles, predefinedRole)
	}
	return roles, nil
}

func setPhotoStudioMemberRoles(
	ctx context.Context,
	txOrDB csql.TxOrDB,
	photoStudioMemberID entity.PhotoStudioMemberID,
	roles []rbac.RoleID,
) ([]*rbac.Role, error) {
	valueStrings := []string{}
	args := []any{}
	for _, roleID := range roles {
		valueStrings = append(valueStrings, "(?,?,NOW())")
		args = append(args, photoStudioMemberID, roleID)
	}
	_, err := csql.ExecContext(
		ctx,
		txOrDB,
		"INSERT IGNORE INTO `photo_studio_member_roles`(`photo_studio_member_id`, `role_id`, `created_at`) VALUES "+
			strings.Join(valueStrings, ","),
		args...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	row := csql.QueryRowContext(
		ctx,
		txOrDB,
		"SELECT COUNT(*) FROM `photo_studio_member_roles` WHERE `photo_studio_member_id` = ?",
		photoStudioMemberID,
	)
	count := -1
	if err := row.Scan(&count); err != nil {
		return nil, terrors.Wrap(err)
	}
	if count > auth.MaxRolesPerPhotoStudioMember {
		return nil, terrors.Wrapf("number of roles is over max : len=%d max=%d", count, auth.MaxRolesPerPhotoStudioMember)
	}
	return getPhotoStudioMemberRoles(ctx, txOrDB, photoStudioMemberID)
}
