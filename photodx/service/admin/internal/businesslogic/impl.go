package businesslogic

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/gateway/line/messaging"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/proc"
)

type Impl struct {
	GenerateChatMessageID  proc.RandomStringGenerator
	Repository             repository.Repository
	LINEMessagingAPIClient messaging.Client
	L                      *slog.Logger
}
