package usecase

import (
	"context"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
)

func (t *Impl) APIPutPushSubscription(
	ctx context.Context,
	principal entity.UserPrincipalRefreshToken,
	pushSubscription *webpush.Subscription,
) (*struct{}, error) {
	return &struct{}{}, t.BusinessLogic.CreateWebPushSubscription(
		ctx,
		pushSubscription,
		principal.GetUserID(),
	)
}
