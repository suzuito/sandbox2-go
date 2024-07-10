package oauth2loginflow

import "time"

type StateCode string

type State struct {
	Code        StateCode
	ProviderID  ProviderID
	RedirectURL string
	CallbackURL string
	ExpiresAt   time.Time
}
