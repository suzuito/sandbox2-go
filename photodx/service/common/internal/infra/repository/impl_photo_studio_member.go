package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func (t *Impl) CreatePhotoStudioMember(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	photoStudioMember *entity.PhotoStudioMember,
	initialPasswordHashValue string,
	initialRoles []rbac.RoleID,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	if _, _, _, err := t.GetPhotoStudioMember(ctx, photoStudioMember.ID); err == nil {
		return nil, nil, nil, terrors.Wrap(&repository.DuplicateEntryError{
			EntryType: repository.EntryTypePhotoStudioMember,
			EntryID:   string(photoStudioMember.ID),
		})
	}
	if err := csql.WithTransaction(ctx, t.Pool, func(tx csql.TxOrDB) error {
		if _, err := csql.ExecContext(
			ctx,
			tx,
			"INSERT INTO `photo_studio_members`(`id`, `photo_studio_id`, `email`, `name`, `active`, `created_at`, `updated_at`) VALUES(?, ?, ?, ?, ?, NOW(), NOW())",
			photoStudioMember.ID,
			photoStudioID,
			photoStudioMember.Email,
			photoStudioMember.Name,
			photoStudioMember.Active,
		); err != nil {
			return terrors.Wrap(err)
		}
		if _, err := csql.ExecContext(
			ctx,
			tx,
			"INSERT INTO `photo_studio_member_password_hash_values`(`photo_studio_member_id`, `value`, `created_at`, `updated_at`) VALUES(?, ?, NOW(), NOW())",
			photoStudioMember.ID,
			initialPasswordHashValue,
		); err != nil {
			return terrors.Wrap(err)
		}
		if _, err := setPhotoStudioMemberRoles(
			ctx,
			tx,
			photoStudioMember.ID,
			initialRoles,
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	return t.GetPhotoStudioMember(ctx, photoStudioMember.ID)
}

func (t *Impl) GetPhotoStudioMember(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id`, `email`, `name`, `active`, `photo_studio_id` FROM `photo_studio_members` WHERE `id` = ?",
		photoStudioMemberID,
	)
	member := entity.PhotoStudioMember{}
	photoStudioID := entity.PhotoStudioID("")
	if err := row.Scan(
		&member.ID,
		&member.Email,
		&member.Name,
		&member.Active,
		&photoStudioID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: repository.EntryTypePhotoStudioMember,
				EntryID:   string(member.ID),
			})
		}
		return nil, nil, nil, terrors.Wrap(err)
	}
	rolesPerPhotoStudioMember, err := getPhotoStudioMemberRoles(ctx, t.Pool, []entity.PhotoStudioMemberID{photoStudioMemberID})
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	photoStudio, err := t.GetPhotoStudio(ctx, photoStudioID)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	return &member, rolesPerPhotoStudioMember[member.ID], photoStudio, nil
}

func (t *Impl) GetPhotoStudioMemberByEmail(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id` FROM `photo_studio_members` WHERE `photo_studio_id` = ? AND `email` = ?",
		photoStudioID,
		email,
	)
	photoStudioMemberID := entity.PhotoStudioMemberID("")
	if err := row.Scan(&photoStudioMemberID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: repository.EntryTypePhotoStudioMember,
				EntryID:   "by email",
			})
		}
		return nil, nil, nil, terrors.Wrap(err)
	}
	return t.GetPhotoStudioMember(ctx, photoStudioMemberID)
}

func (t *Impl) GetPhotoStudioMemberPasswordHashByEmail(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
) (string, *entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	member, roles, photoStudio, err := t.GetPhotoStudioMemberByEmail(ctx, photoStudioID, email)
	if err != nil {
		return "", nil, nil, nil, terrors.Wrap(err)
	}
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `value` FROM `photo_studio_member_password_hash_values` WHERE `photo_studio_member_id` = ?",
		member.ID,
	)
	hashValue := ""
	if err := row.Scan(&hashValue); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil, nil, nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: repository.EntryTypePhotoStudioMember,
				EntryID:   string(member.ID),
			})
		}
		return "", nil, nil, nil, terrors.Wrap(err)
	}
	return hashValue, member, roles, photoStudio, nil
}
