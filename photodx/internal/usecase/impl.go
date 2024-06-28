package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service"
)

type Impl struct {
	S service.Service
	L *slog.Logger
}
