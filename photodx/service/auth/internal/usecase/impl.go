package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
)

type Impl struct {
	B businesslogic.BusinessLogic
	L *slog.Logger
}
