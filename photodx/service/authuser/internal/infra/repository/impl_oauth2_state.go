package repository

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

func (t *Impl) CreateOAuth2State(
	ctx context.Context,
	state *oauth2loginflow.State,
) (*oauth2loginflow.State, error) {
	mState := newModelOAuth2State(state)
	if err := t.GormDB.Create(mState).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mState.ToEntity(), nil
}

func (t *Impl) GetAndDeleteOAuth2State(
	ctx context.Context,
	stateCode oauth2loginflow.StateCode,
) (*oauth2loginflow.State, error) {
	mState := modelOAuth2State{}
	if err := t.GormDB.
		WithContext(ctx).
		Where("code = ?", stateCode).
		First(&mState).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &repository.NoEntryError{
				EntryType: "OAuth2State",
				EntryID:   string(stateCode),
			}
		}
	}
	if err := t.GormDB.
		WithContext(ctx).
		Where("code = ?", stateCode).
		Delete(&modelOAuth2State{}).
		Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mState.ToEntity(), nil
}
