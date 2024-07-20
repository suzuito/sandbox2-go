package entity

import (
	"time"

	"github.com/SherClockHolmes/webpush-go"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type UserWebPushSubscription struct {
	Endpoint       string
	UserID         common_entity.UserID
	ExpirationTime time.Time
	Value          *webpush.Subscription
}
