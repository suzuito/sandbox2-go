package usecase

import (
	"log/slog"

	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/businesslogic"
)

type Impl struct {
	BusinessLogic       businesslogic.BusinessLogic
	CommonBusinessLogic common_businesslogic.BusinessLogic
	L                   *slog.Logger
}
