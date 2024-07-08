package businesslogic

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
)

type Impl struct {
	L                           *slog.Logger
	AdminAccessTokenJWTVerifier auth.JWTVerifier
}
