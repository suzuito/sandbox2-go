package usecase

import (
	"context"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIPostUserCreationRequest struct {
	UserCreationRequestID common_entity.UserCreationRequestID `json:"userCreationRequestId"`
}

func (t *Impl) APIPostUserCreationRequest(
	ctx context.Context,
	frontBaseURL *url.URL,
	email string,
) (*DTOAPIPostUserCreationRequest, error) {
	req, err := t.BusinessLogic.CreateUserCreationRequest(
		ctx,
		email,
		600,
		frontBaseURL,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostUserCreationRequest{
		UserCreationRequestID: req.ID,
	}, nil
}

type DTOAPIPostUserCreation struct {
}

func (t *Impl) APIPostUserCreation(
	ctx context.Context,
	userCreationRequestID common_entity.UserCreationRequestID,
) (*DTOAPIPostUserCreation, error) {
	_, err := t.BusinessLogic.GetUserCreationRequestNotExpired(ctx, userCreationRequestID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostUserCreation{}, nil
}

type DTOAPIPostUserCreationVerifyResult string

const (
	DTOAPIPostUserCreationVerifyResultOK DTOAPIPostUserCreationVerifyResult = "ok"
	DTOAPIPostUserCreationVerifyResultNG DTOAPIPostUserCreationVerifyResult = "ng"
)

type DTOAPIPostUserCreationVerify struct {
}

func (t *Impl) APIPostUserCreationVerify(
	ctx context.Context,
	userCreationRequestID common_entity.UserCreationRequestID,
	code common_entity.UserCreationCode,
) (*DTOAPIPostUserCreationVerify, error) {
	if _, err := t.BusinessLogic.GetValidUserCreationRequestNotExpired(
		ctx,
		userCreationRequestID,
		code,
	); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostUserCreationVerify{}, nil
}

type DTOAPIPostUserCreationCreate struct {
	User         *common_entity.User `json:"user"`
	RefreshToken string              `json:"refreshToken"`
}

func (t *Impl) APIPostUserCreationCreate(
	ctx context.Context,
	userCreationRequestID common_entity.UserCreationRequestID,
	code common_entity.UserCreationCode,
	plainPassword string,
) (*DTOAPIPostUserCreationCreate, error) {
	req, err := t.BusinessLogic.GetValidUserCreationRequestNotExpired(
		ctx,
		userCreationRequestID,
		code,
	)
	user, err := t.BusinessLogic.CreateUser(
		ctx,
		req.Email,
		plainPassword,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.BusinessLogic.DeleteUserCreationRequest(ctx, userCreationRequestID); err != nil {
		return nil, terrors.Wrap(err)
	}
	refreshToken, err := t.BusinessLogic.CreateUserRefreshToken(ctx, user.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostUserCreationCreate{
		User:         user,
		RefreshToken: refreshToken,
	}, nil
}
