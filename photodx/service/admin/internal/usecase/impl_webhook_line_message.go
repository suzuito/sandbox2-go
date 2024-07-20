package usecase

import (
	"context"
	"fmt"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
)

func (t *Impl) procLINEMessagingAPIWebhook_Message(
	ctx context.Context,
	lineLinkInfo *entity.LineLinkInfo,
	messageBytes []byte,
) error {
	return fmt.Errorf("not impl")
}
