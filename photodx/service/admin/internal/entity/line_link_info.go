package entity

import "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"

type LineLinkInfo struct {
	PhotoStudioID             entity.PhotoStudioID `json:"photoStudioId"`
	MessagingAPIChannelSecret string               `json:"messagingApiChannelSecret"`
	LongAccessToken           string               `json:"longAccessToken"`
	Active                    bool                 `json:"active"`
}
