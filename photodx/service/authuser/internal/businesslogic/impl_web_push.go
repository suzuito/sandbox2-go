package businesslogic

import (
	"context"
	"log/slog"
	"time"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) GetWebPushVAPIDPublicKey(
	ctx context.Context,
	userID common_entity.UserID,
) (string, error) {
	return t.WebPushVAPIDPublicKey, nil
}

func (t *Impl) CreateWebPushSubscription(
	ctx context.Context,
	subscription *webpush.Subscription,
	userID common_entity.UserID,
) error {
	_, err := t.Repository.UpdateOrCreateUserWebPushSubscription(
		ctx,
		&entity.UserWebPushSubscription{
			Endpoint:       subscription.Endpoint,
			UserID:         userID,
			ExpirationTime: t.NowFunc().Add(3 * 24 * time.Hour),
			Value:          subscription,
		},
	)
	return terrors.Wrap(err)
}

func (t *Impl) PushNotification(
	ctx context.Context,
	l *slog.Logger,
	userID common_entity.UserID,
	message string,
) error {
	pushSubscriptions, err := t.Repository.GetLatestUserWebPushSubscriptions(ctx, userID)
	if err != nil {
		return terrors.Wrap(err)
	}
	if len(pushSubscriptions) <= 0 {
		return terrors.Wrap(err)
	}
	for _, s := range pushSubscriptions {
		_, err := webpush.SendNotificationWithContext(
			ctx,
			[]byte(message),
			s.Value,
			&webpush.Options{
				VAPIDPublicKey:  t.WebPushVAPIDPublicKey,
				VAPIDPrivateKey: t.WebPushVAPIDPrivateKey,
			},
		)
		if err != nil {
			l.Warn("failed to webpush.SendNotificationWithContext", "err", err)
			continue
		}
	}
	return nil
}
