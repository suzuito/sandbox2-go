package cwebpush

import (
	"context"
	"io"
	"net/http"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func SendNotificationWithContext(
	ctx context.Context,
	vapidPublicKey string,
	vapidPrivateKey string,
	data []byte,
	sub *webpush.Subscription,
) error {
	res, err := webpush.SendNotificationWithContext(
		ctx,
		data,
		sub,
		&webpush.Options{
			VAPIDPublicKey:  vapidPublicKey,
			VAPIDPrivateKey: vapidPrivateKey,
			// SafariのWebPushにて、subの指定が必須となっている。
			// https://developer.apple.com/documentation/usernotifications/sending-web-push-notifications-in-web-apps-and-browsers#:~:text=to%20the%20specification.-,BadJwtToken,-The%20JSON%20web
			Subscriber: "example@example.com",
		},
	)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return terrors.Wrapf("cannot web push by http error: endpoint=%s status=%d body=%s", sub.Endpoint, res.StatusCode, string(body))
	}
	return nil
}
