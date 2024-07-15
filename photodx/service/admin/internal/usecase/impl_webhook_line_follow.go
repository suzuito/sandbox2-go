package usecase

import (
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) procLINEMessagingAPIWebhook_Follow(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	messageBytes []byte,
) error {
	message := entity.LINEWebhookEventFollow{}
	if err := json.Unmarshal(messageBytes, &message); err != nil {
		return terrors.Wrap(err)
	}
	// Create user
	if message.Source != nil {
		_, err := t.AuthUserBusinessLogic.CreateUserIfNotExistsFromResourceOwnerID(
			ctx,
			"line",
			message.Source.UserID,
		)
		if err != nil {
			t.L.Error("CreateUserIfNotExistsFromResourceOwnerID is failed", "err", err)
		}
	}
	return nil
}
