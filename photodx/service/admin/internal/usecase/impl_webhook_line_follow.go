package usecase

import (
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
)

func (t *Impl) procLINEMessagingAPIWebhook_Follow(
	ctx context.Context,
	lineLinkInfo *entity.LineLinkInfo,
	messageBytes []byte,
) error {
	message := entity.LINEWebhookEventFollow{}
	if err := json.Unmarshal(messageBytes, &message); err != nil {
		return terrors.Wrap(err)
	}
	if message.Source == nil {
		return terrors.Wrapf("message.Source is not found in message : %+v", message)
	}
	if message.Source.Type != "user" {
		return terrors.Wrapf("message.Source.Type is not user : %+v", message)
	}
	user, err := t.BusinessLogic.GenerateUserFromLINEProfile(ctx, lineLinkInfo, message.Source.UserID)
	if err != nil {
		return terrors.Wrap(err)
	}
	// Create user
	if message.Source != nil {
		created, err := t.AuthUserBusinessLogic.CreateUserIfNotExistsFromResourceOwnerID(
			ctx,
			"line",
			message.Source.UserID,
			user,
		)
		if err != nil {
			return terrors.Wrap(err)
		}
		if _, err := t.BusinessLogic.CreatePhotoStudioUser(ctx, lineLinkInfo.PhotoStudioID, created); err != nil {
			return terrors.Wrap(err)
		}
	}
	return nil
}
