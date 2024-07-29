package usecase

import (
	"log/slog"
	"time"

	admin_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/admin/pkg/businesslogic"
	auth_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/auth/pkg/businesslogic"
	authuser_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/businesslogic"
	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/businesslogic"
)

type Impl struct {
	NowFunc               func() time.Time
	BusinessLogic         businesslogic.BusinessLogic
	CommonBusinessLogic   common_businesslogic.BusinessLogic
	AuthBusinessLogic     auth_businesslogic.ExposedBusinessLogic
	AuthUserBusinessLogic authuser_businesslogic.ExposedBusinessLogic
	AdminBusinessLogic    admin_businesslogic.ExposedBusinessLogic
	L                     *slog.Logger
}
