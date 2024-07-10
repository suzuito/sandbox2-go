package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
)

type Impl struct {
	BusinessLogic       businesslogic.BusinessLogic
	CommonBusinessLogic common_businesslogic.BusinessLogic
	L                   *slog.Logger

	OAuth2ProviderLINE *oauth2loginflow.Provider
}
