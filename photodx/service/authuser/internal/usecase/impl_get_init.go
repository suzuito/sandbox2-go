package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIGetInit struct {
	User                  *common_entity.User `json:"user"`
	WebPushVAPIDPublicKey string              `json:"vapidPublicKey"`
}

func (t *Impl) APIGetInit(ctx context.Context, principal common_entity.UserPrincipalAccessToken) (*DTOAPIGetInit, error) {
	user, err := t.BusinessLogic.GetUser(ctx, principal.GetUserID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	vapidPublicKey, err := t.BusinessLogic.GetWebPushVAPIDPublicKey(ctx, principal.GetUserID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetInit{
		User:                  user,
		WebPushVAPIDPublicKey: vapidPublicKey,
	}, nil
}
