package entity

import (
	"time"

	"github.com/SherClockHolmes/webpush-go"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type PhotoStudioMemberWebPushSubscription struct {
	Endpoint            string
	PhotoStudioMemberID common_entity.PhotoStudioMemberID
	ExpirationTime      time.Time
	Value               *webpush.Subscription
}
