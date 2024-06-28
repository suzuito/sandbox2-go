package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) GetPhotoStudioMemberPasswordHashByEmail(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	email string,
) (string, error) {
	member, err := t.GetPhotoStudioMemberByEmail(ctx, photoStudioID, email)
	if err != nil {
		return "", terrors.Wrap(err)
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
			return "", terrors.Wrap(&repository.NoEntryError{
				EntryType: "PhotoStudioMemberHashValues",
			})
		}
		return "", terrors.Wrap(err)
	}
	return hashValue, nil
}
