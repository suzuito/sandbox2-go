package usecase

import (
	"context"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIPostRequestPromoteGuestUser struct {
}

func (t *Impl) APIPostRequestPromoteGuestUser(
	ctx context.Context,
	frontBaseURL url.URL,
	userID common_entity.UserID,
	email string,
) (*DTOAPIPostRequestPromoteGuestUser, error) {
	if err := t.BusinessLogic.RequestPromoteGuestUser(
		ctx,
		frontBaseURL,
		userID,
		email,
		600,
	); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostRequestPromoteGuestUser{}, nil
}

type DTOAPIPostPromoteGuestUser struct {
	User         *common_entity.User `json:"user"`
	RefreshToken string              `json:"refreshToken"`
}

func (t *Impl) APIPostApprovePromoteGuestUser(
	ctx context.Context,
	principal common_entity.UserPrincipalAccessToken,
	plainPassword string,
	code string,
) (*DTOAPIPostPromoteGuestUser, error) {
	user, err := t.BusinessLogic.PromoteGuestUser(
		ctx,
		principal.GetUserID(),
		plainPassword,
		code,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.BusinessLogic.CreateUserRefreshToken(ctx, user.ID, user.Guest)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostPromoteGuestUser{
		User:         user,
		RefreshToken: refreshToken,
	}, nil
}
