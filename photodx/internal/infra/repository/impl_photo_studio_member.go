package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/auth/predefined"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) CreatePhotoStudioMember(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	photoStudioMember *entity.PhotoStudioMember,
	initialPasswordHashValue string,
) (*entity.PhotoStudioMember, error) {
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
			[]rbac.RoleID{
				predefined.RoleSuperUser.ID, // TODO スーパーユーザー権限を与えているのはデバッグ用途のため。後で変えてください。
			},
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, terrors.Wrap(err)
	}
	created, err := t.GetPhotoStudioMember(ctx, photoStudioMember.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}

func (t *Impl) getPhotoStudioMember(
	ctx context.Context,
	cond string,
	condArgs []any,
) (*entity.PhotoStudioMember, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id`, `email`, `name`, `active` FROM `photo_studio_members`"+cond,
		condArgs...,
	)
	member := entity.PhotoStudioMember{}
	if err := row.Scan(
		&member.ID,
		&member.Email,
		&member.Name,
		&member.Active,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: "PhotoStudioMember",
			})
		}
		return nil, terrors.Wrap(err)
	}
	return &member, nil
}

func (t *Impl) GetPhotoStudioMember(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (*entity.PhotoStudioMember, error) {
	return t.getPhotoStudioMember(
		ctx,
		"WHERE `id` = ?",
		[]any{photoStudioMemberID},
	)
}

func (t *Impl) GetPhotoStudioMemberByEmail(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
) (*entity.PhotoStudioMember, error) {
	return t.getPhotoStudioMember(
		ctx,
		"WHERE `photo_studio_id` = ? AND `email` = ?",
		[]any{photoStudioID, email},
	)
}
