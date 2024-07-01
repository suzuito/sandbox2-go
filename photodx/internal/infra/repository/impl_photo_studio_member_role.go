package repository

import (
	"context"
	"strings"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

func getPhotoStudioMemberRoles(
	ctx context.Context,
	txOrDB csql.TxOrDB,
	photoStudioMemberIDs []entity.PhotoStudioMemberID,
) (map[entity.PhotoStudioMemberID][]*rbac.Role, error) {
	rows, err := csql.QueryContext(
		ctx,
		txOrDB,
		"SELECT `role_id`, `photo_studio_member_id` FROM `photo_studio_member_roles` WHERE "+csql.SqlIn("photo_studio_member_id", photoStudioMemberIDs),
		csql.ToAnySlice(photoStudioMemberIDs)...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rows.Close()
	rolesPerPhotoStudioMember := map[entity.PhotoStudioMemberID][]*rbac.Role{}
	for rows.Next() {
		roleID := rbac.RoleID("")
		photoStudioMemberID := entity.PhotoStudioMemberID("")
		if err := rows.Scan(
			&roleID,
			&photoStudioMemberID,
		); err != nil {
			return nil, terrors.Wrap(err)
		}
		role, exists := rbac.AvailablePredefinedRoles[roleID]
		if !exists {
			continue
		}
		roles := rolesPerPhotoStudioMember[photoStudioMemberID]
		roles = append(roles, role)
		rolesPerPhotoStudioMember[photoStudioMemberID] = roles
	}
	return rolesPerPhotoStudioMember, nil
}

func setPhotoStudioMemberRoles(
	ctx context.Context,
	txOrDB csql.TxOrDB,
	photoStudioMemberID entity.PhotoStudioMemberID,
	roleIDs []rbac.RoleID,
) ([]*rbac.Role, error) {
	valueStrings := []string{}
	args := []any{}
	for _, roleID := range roleIDs {
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
	if count > entity.MaxRolesPerPhotoStudioMember {
		return nil, terrors.Wrapf("number of roles is over max : len=%d max=%d", count, entity.MaxRolesPerPhotoStudioMember)
	}
	rolesPerPhotoStudioMember, err := getPhotoStudioMemberRoles(ctx, txOrDB, []entity.PhotoStudioMemberID{photoStudioMemberID})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	rolesCreated := rolesPerPhotoStudioMember[photoStudioMemberID]
	return rolesCreated, nil
}
