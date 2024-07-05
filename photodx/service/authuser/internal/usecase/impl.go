package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
)

type Impl struct {
	B businesslogic.BusinessLogic
	L *slog.Logger

	OAuth2ProviderLINE *oauth2loginflow.Provider
}
