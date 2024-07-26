package usecase

import (
	"log/slog"
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/businesslogic"
	auth_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/auth/pkg/businesslogic"
	authuser_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/businesslogic"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
)

type Impl struct {
	NowFunc               func() time.Time
	BusinessLogic         businesslogic.BusinessLogic
	CommonBusinessLogic   common_businesslogic.BusinessLogic
	AuthUserBusinessLogic authuser_businesslogic.ExposedBusinessLogic
	AuthBusinessLogic     auth_businesslogic.ExposedBusinessLogic
	L                     *slog.Logger
}
