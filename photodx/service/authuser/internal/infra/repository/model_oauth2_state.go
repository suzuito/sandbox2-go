package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
)

type modelOAuth2State struct {
	Code        oauth2loginflow.StateCode
	ProviderID  oauth2loginflow.ProviderID
	RedirectURL string
	ExpiresAt   time.Time
}

func (t *modelOAuth2State) TableName() string {
	return "oauth2_loginflow_states"
}

func (t *modelOAuth2State) ToEntity() *oauth2loginflow.State {
	return &oauth2loginflow.State{
		Code:        t.Code,
		ProviderID:  t.ProviderID,
		RedirectURL: t.RedirectURL,
		ExpiresAt:   t.ExpiresAt,
	}
}

func newModelOAuth2State(s *oauth2loginflow.State) *modelOAuth2State {
	return &modelOAuth2State{
		Code:        s.Code,
		ProviderID:  s.ProviderID,
		RedirectURL: s.RedirectURL,
		ExpiresAt:   s.ExpiresAt,
	}
}
