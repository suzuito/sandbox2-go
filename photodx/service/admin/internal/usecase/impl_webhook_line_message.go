package usecase

import (
	"context"
	"fmt"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) procLINEMessagingAPIWebhook_Message(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	messageBytes []byte,
) error {
	return fmt.Errorf("not impl")
}
