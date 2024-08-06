package businesslogic

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func (t *Impl) CreateUserCreationRequest(
	ctx context.Context,
	email string,
	ttlSeconds int,
	frontURL *url.URL,
) (*common_entity.UserCreationRequest, error) {
	var noEntryError *repository.NoEntryError
	if _, err := t.Repository.GetUserByEmail(ctx, email); err == nil {
		return nil, terrors.Wrap(
			&common_repository.DuplicateEntryError{
				EntryType: "User",
				EntryID:   fmt.Sprintf("email:%s", email),
			},
		)
	} else if !errors.As(err, &noEntryError) {
		return nil, terrors.Wrap(err)
	}
	id, err := t.UserCreationRequestIDGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	code, err := t.UserCreationCodeGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	now := t.NowFunc()
	req := common_entity.UserCreationRequest{
		ID:        common_entity.UserCreationRequestID(id),
		Email:     email,
		Code:      common_entity.UserCreationCode(code),
		ExpiredAt: now.Add(time.Duration(ttlSeconds) * time.Second),
	}
	if err := t.Repository.CreateUserCreationRequest(ctx, &req); err != nil {
		return nil, terrors.Wrap(err)
	}
	userRegisterURL := *frontURL
	userRegisterURL.Path = "/register"
	if err := t.UserMailSender.SendUserCreationCode(ctx, req, &userRegisterURL); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &req, nil
}

func (t *Impl) GetValidUserCreationRequestNotExpired(
	ctx context.Context,
	id common_entity.UserCreationRequestID,
	code common_entity.UserCreationCode,
) (*common_entity.UserCreationRequest, error) {
	now := t.NowFunc()
	r, err := t.Repository.GetUserCreationRequest(
		ctx,
		id,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if now.After(r.ExpiredAt) {
		return nil, &common_repository.NoEntryError{
			EntryType: "User",
			EntryID:   string(id),
		}
	}
	if r.Code != code {
		return nil, terrors.Wrap(ErrMismtachUserCreationRequestCode)
	}
	return r, nil
}

func (t *Impl) CreateUser(
	ctx context.Context,
	email string,
	planinPassword string,
) (*common_entity.User, error) {
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
		Email:           email,
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
