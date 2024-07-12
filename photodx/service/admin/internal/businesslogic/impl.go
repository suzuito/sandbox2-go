package businesslogic

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
)

type Impl struct {
	Repository repository.Repository
	L          *slog.Logger
}
