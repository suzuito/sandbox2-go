package usecase

import (
	"context"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAuthPutPushSubscription struct{}

func (t *Impl) AuthPutPushSubscription(
	ctx context.Context,
	principal entity.AdminPrincipalAccessToken,
	sub *webpush.Subscription,
) (*DTOAuthPutPushSubscription, error) {
	if err := t.BusinessLogic.CreateWebPushSubscription(ctx, sub, principal.GetPhotoStudioMemberID()); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAuthPutPushSubscription{}, nil
}
