package businesslogic

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/cwebpush"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) GetWebPushVAPIDPublicKey(
	ctx context.Context,
	userID common_entity.PhotoStudioMemberID,
) (string, error) {
	return t.WebPushVAPIDPublicKey, nil
}

func (t *Impl) CreateWebPushSubscription(
	ctx context.Context,
	subscription *webpush.Subscription,
	photoStudioMemberID common_entity.PhotoStudioMemberID,
) error {
	_, err := t.Repository.UpdateOrCreateUserWebPushSubscription(
		ctx,
		&entity.PhotoStudioMemberWebPushSubscription{
			Endpoint:            subscription.Endpoint,
			PhotoStudioMemberID: photoStudioMemberID,
			ExpirationTime:      t.NowFunc().Add(3 * 24 * time.Hour),
			Value:               subscription,
		},
	)
	return terrors.Wrap(err)
}

func (t *Impl) PushNotification(
	ctx context.Context,
	l *slog.Logger,
	photoStudioMemberID common_entity.PhotoStudioMemberID,
	notification *common_entity.Notification,
) error {
	if notification == nil {
		return terrors.Wrapf("notification is nil")
	}
	notificationBytes, err := json.Marshal(notification)
	if err != nil {
		return terrors.Wrap(err)
	}
	pushSubscriptions, err := t.Repository.GetLatestPhotoStudioMemberWebPushSubscriptions(ctx, photoStudioMemberID)
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

func (t *Impl) PushNotificationToAllMembers(
	ctx context.Context,
	l *slog.Logger,
	photoStudioID common_entity.PhotoStudioID,
	notification *common_entity.Notification,
) error {
	listQuery := cgorm.ListQuery{
		Offset: 0,
		Limit:  10,
	}
	for {
		members, hasNext, err := t.Repository.ListPhotoStudioMembers(ctx, photoStudioID, &listQuery)
		if err != nil {
			return terrors.Wrap(err)
		}
		for _, member := range members {
			t.PushNotification(ctx, l, member.ID, notification)
		}
		if !hasNext {
			break
		}
		listQuery.Offset = listQuery.NextOffset()
	}
	return nil
}
