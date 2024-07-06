package repository

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

func (t *Impl) CreatePhotoStudioMember(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	photoStudioMember *entity.PhotoStudioMember,
	initialPasswordHashValue string,
	initialRoles []rbac.RoleID,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	if err := t.GormDB.
		WithContext(ctx).
		Where(photoStudioMember.ID).
		First(&modelPhotoStudioMember{}).Error; err == nil {
		return nil, nil, nil, terrors.Wrap(&repository.DuplicateEntryError{
			EntryType: repository.EntryTypePhotoStudioMember,
			EntryID:   string(photoStudioMember.ID),
		})
	}
	if err := t.GormDB.Transaction(func(tx *gorm.DB) error {
		mPhotoStudioMember := newModelPhotoStudioMember(photoStudioMember)
		mPhotoStudioMember.CreatedAt = t.NowFunc()
		mPhotoStudioMember.UpdatedAt = t.NowFunc()
		if err := tx.Create(mPhotoStudioMember).Error; err != nil {
			return terrors.Wrap(err)
		}
		mPasswordHash := modelPhotoStudioMemberPasswordHashValue{
			PhotoStudioMemberID: mPhotoStudioMember.ID,
			Value:               initialPasswordHashValue,
			CreatedAt:           t.NowFunc(),
			UpdatedAt:           t.NowFunc(),
		}
		if err := tx.Create(&mPasswordHash).Error; err != nil {
			return terrors.Wrap(err)
		}
		mRoles := modelPhotoStudioMemberRoles{}
		for _, initialRole := range initialRoles {
			mRoles = append(
				mRoles,
				&modelPhotoStudioMemberRole{
					PhotoStudioMemberID: mPhotoStudioMember.ID,
					RoleID:              initialRole,
					CreatedAt:           t.NowFunc(),
				},
			)
		}
		if err := tx.Create(&mRoles).Error; err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	return t.GetPhotoStudioMember(ctx, photoStudioMember.ID)
}

/*
func (t *Impl) CreatePhotoStudioMemberOld(
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
*/

func (t *Impl) GetPhotoStudioMember(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	return t.getPhotoStudioMember(ctx, []gormQueryWhere{
		{query: photoStudioMemberID},
	})
}

func (t *Impl) GetPhotoStudioMemberByEmail(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	return t.getPhotoStudioMember(ctx, []gormQueryWhere{
		{query: "photo_studio_id = ?", args: []any{photoStudioID}},
		{query: "email = ?", args: []any{email}},
	})
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
	passwordHash := modelPhotoStudioMemberPasswordHashValue{}
	if err := t.GormDB.
		WithContext(ctx).
		Where("photo_studio_member_id = ?", member.ID).
		First(&passwordHash).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, nil, nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: "PhotoStudioMemberPasswordHash",
				EntryID:   string(member.ID),
			})
		}
		return "", nil, nil, nil, terrors.Wrap(err)
	}
	return passwordHash.Value, member, roles, photoStudio, nil
}

func (t *Impl) getPhotoStudioMember(
	ctx context.Context,
	wheres []gormQueryWhere,
) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error) {
	mPhotoStudioMember := modelPhotoStudioMember{}
	db := t.GormDB.WithContext(ctx)
	for _, where := range wheres {
		db = db.Where(where.query, where.args...)
	}
	db = db.
		Preload("PhotoStudio").
		Preload("Roles")
	if err := db.First(&mPhotoStudioMember).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, &repository.NoEntryError{
				EntryType: repository.EntryTypePhotoStudioMember,
				EntryID:   string(mPhotoStudioMember.ID),
			}
		}
		return nil, nil, nil, terrors.Wrap(err)
	}
	return mPhotoStudioMember.ToEntity(),
		mPhotoStudioMember.Roles.ToEntity(),
		mPhotoStudioMember.PhotoStudio.ToEntity(),
		nil
}
