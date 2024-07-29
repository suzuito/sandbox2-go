package businesslogic

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/common/cwebpush"
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
	notification *common_entity.Notification,
) error {
	if notification == nil {
		return terrors.Wrapf("notification is nil")
	}
	notificationBytes, err := json.Marshal(notification)
	if err != nil {
		return terrors.Wrap(err)
	}
	pushSubscriptions, err := t.Repository.GetLatestUserWebPushSubscriptions(ctx, userID)
	if err != nil {
		return terrors.Wrap(err)
	}
	if len(pushSubscriptions) <= 0 {
		return nil
	}
	for _, s := range pushSubscriptions {
		if err := cwebpush.SendNotificationWithContext(ctx, t.WebPushVAPIDPublicKey, t.WebPushVAPIDPrivateKey, notificationBytes, s.Value); err != nil {
			l.Warn("failed to webpush.SendNotificationWithContext", "err", err)
			continue
		}
	}
	return nil
}
