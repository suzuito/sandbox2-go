package usecase

import (
	"context"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	pkg_entity "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIPostRegisterRequest struct {
	Code pkg_entity.RequestRegisterUserResultCode `json:"code"`
}

func (t *Impl) APIPostRegisterRequest(
	ctx context.Context,
	frontBaseURL *url.URL,
	email string,
) (*DTOAPIPostRegisterRequest, error) {
	result, err := t.BusinessLogic.RequestRegisterUser(
		ctx,
		email,
		600,
		frontBaseURL,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostRegisterRequest{
		Code: result,
	}, nil
}

type DTOAPIPostRegisterApprove struct {
	User         *common_entity.User `json:"user"`
	RefreshToken string              `json:"refreshToken"`
}

func (t *Impl) APIPostRegisterApprove(
	ctx context.Context,
	userCreationRequestID common_entity.UserCreationRequestID,
	code common_entity.UserCreationCode,
	plainPassword string,
) (*DTOAPIPostRegisterApprove, error) {
	user, err := t.BusinessLogic.RegisterUser(
		ctx,
		userCreationRequestID,
		code,
		plainPassword,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.BusinessLogic.CreateUserRefreshToken(ctx, user.ID, user.Guest)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostRegisterApprove{
		User:         user,
		RefreshToken: refreshToken,
	}, nil
}
