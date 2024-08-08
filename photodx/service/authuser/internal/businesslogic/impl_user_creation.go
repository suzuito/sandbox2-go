package businesslogic

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
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
	if _, err := t.Repository.GetUserCreationRequestByEmail(ctx, email); err == nil {
		return nil, terrors.Wrap(
			&common_repository.DuplicateEntryError{
				EntryType: "UserCreationRequest",
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
	r, err := t.GetUserCreationRequestNotExpired(ctx, id)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if r.Code != code {
		return nil, terrors.Wrap(ErrMismatchUserCreationCode)
	}
	return r, nil
}

func (t *Impl) GetUserCreationRequestNotExpired(
	ctx context.Context,
	id common_entity.UserCreationRequestID,
) (*common_entity.UserCreationRequest, error) {
	r, err := t.Repository.GetUserCreationRequest(
		ctx,
		id,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	now := t.NowFunc()
	if now.After(r.ExpiredAt) {
		return nil, &common_repository.NoEntryError{
			EntryType: "User",
			EntryID:   string(id),
		}
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
		Name:            fmt.Sprintf("ユーザー%d%d", time.Now().Unix(), rand.Intn(1000)),
		Email:           email,
		EmailVerified:   true,
		ProfileImageURL: "https://vos.line-scdn.net/chdev-console-static/default-profile.png",
		Active:          true,
	}
	if err := user.Validate(); err != nil {
		return nil, terrors.Wrap(err)
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

func (t *Impl) DeleteUserCreationRequest(
	ctx context.Context,
	userCreationRequestID common_entity.UserCreationRequestID,
) error {
	return terrors.Wrap(t.Repository.DeleteUserCreationRequest(ctx, userCreationRequestID))
}
