package usecase

import (
	"context"

	webpush "github.com/SherClockHolmes/webpush-go"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) APIPutPushSubscription(
	ctx context.Context,
	principal common_entity.UserPrincipalAccessToken,
	pushSubscription *webpush.Subscription,
) (*struct{}, error) {
	return &struct{}{}, t.BusinessLogic.CreateWebPushSubscription(
		ctx,
		pushSubscription,
		principal.GetUserID(),
	)
}
