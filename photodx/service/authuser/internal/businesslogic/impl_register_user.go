package businesslogic

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	pkg_entity "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func (t *Impl) RequestRegisterUser(
	ctx context.Context,
	email string,
	ttlSeconds int,
	frontURL *url.URL,
) (pkg_entity.RequestRegisterUserResultCode, error) {
	var noEntryError *repository.NoEntryError
	if _, err := t.Repository.GetUserByEmail(ctx, email); err == nil {
		return pkg_entity.RequestRegisterUserResultCodeEmailAlreadyExisted, nil
	} else if !errors.As(err, &noEntryError) {
		return -1, terrors.Wrap(err)
	}
	id, err := t.UserCreationRequestIDGenerator.Gen()
	if err != nil {
		return -1, terrors.Wrap(err)
	}
	code, err := t.UserCreationCodeGenerator.Gen()
	if err != nil {
		return -1, terrors.Wrap(err)
	}
	now := t.NowFunc()
	req := common_entity.UserCreationRequest{
		ID:        common_entity.UserCreationRequestID(id),
		Email:     email,
		Code:      common_entity.UserCreationCode(code),
		ExpiredAt: now.Add(time.Duration(ttlSeconds) * time.Second),
	}
	if err := t.Repository.CreateUserCreationRequest(ctx, &req); err != nil {
		return -1, terrors.Wrap(err)
	}
	userRegisterURL := *frontURL
	userRegisterURL.Path = "/register"
	if err := t.UserMailSender.SendUserCreationCode(ctx, req, &userRegisterURL); err != nil {
		return -1, terrors.Wrap(err)
	}
	return pkg_entity.RequestRegisterUserResultCodeCreated, nil
}

func (t *Impl) RegisterUser(
	ctx context.Context,
	id common_entity.UserCreationRequestID,
	code common_entity.UserCreationCode,
	planinPassword string,
) (*common_entity.User, error) {
	now := t.NowFunc()
	req, err := t.Repository.GetUserCreationRequest(
		ctx,
		id,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if now.After(req.ExpiredAt) {
		return nil, terrors.Wrap(&common_repository.NoEntryError{
			EntryType: "UserCreationRequest",
			EntryID:   string(id),
		})
	}
	if req.Code != code {
		return nil, terrors.Wrapf("code is mismatch")
	}
	hashedPassword := t.PasswordHasher.Gen(
		[]byte(t.PasswordSalt),
		planinPassword,
	)
	userID, err := t.UserIDGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	user := common_entity.User{
		ID:              common_entity.UserID(userID),
		Name:            "",
		Email:           req.Email,
		EmailVerified:   true,
		ProfileImageURL: "",
		Active:          true,
		Guest:           false,
	}
	created, err := t.Repository.CreateUser(
		ctx,
		&user,
		hashedPassword,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}
