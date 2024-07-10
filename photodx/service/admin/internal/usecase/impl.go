package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
)

type Impl struct {
	BusinessLogic       businesslogic.BusinessLogic
	CommonBusinessLogic common_businesslogic.BusinessLogic
	L                   *slog.Logger
}
