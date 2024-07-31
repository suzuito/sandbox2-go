package businesslogic

import (
	"context"
	"fmt"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) RequestPromoteGuestUser(
	ctx context.Context,
	frontURLBase url.URL,
	userID common_entity.UserID,
	email string,
	ttlSeconds int,
) error {
	user, err := t.Repository.GetUser(ctx, userID)
	if err != nil {
		return terrors.Wrap(err)
	}
	if !user.Guest {
		return terrors.Wrapf("this user is not guest")
	}
	code, err := t.PromoteGuestUserConfirmationCodeGenerator.Gen()
	if err != nil {
		return terrors.Wrap(err)
	}
	if err := t.Repository.CreatePromoteGuestUserConfirmationCode(
		ctx,
		userID,
		email,
		code,
		ttlSeconds,
	); err != nil {
		return terrors.Wrap(err)
	}
	frontURLBase.Path = fmt.Sprintf("/promote/%s", code)
	if err := t.UserMailSender.SendUserCreationMail(ctx, email, &frontURLBase); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) PromoteGuestUser(
	ctx context.Context,
	userID common_entity.UserID,
	planinPassword string,
	code string,
) (*common_entity.User, error) {
	userConfirmationCode, err := t.Repository.GetPromoteGuestUserConfirmationCodeNotExpired(
		ctx,
		userID,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if userConfirmationCode.Code != code {
		return nil, terrors.Wrapf("code is mismatch")
	}
	hashedPassword := t.PasswordHasher.Gen(
		[]byte(t.PasswordSalt),
		planinPassword,
	)
	user, err := t.Repository.PromoteUser(
		ctx,
		userID,
		userConfirmationCode.Email,
		true,
		hashedPassword,
		true,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return user, nil
}
