package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	admin_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/admin/pkg/businesslogic"
	auth_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/auth/pkg/businesslogic"
	authuser_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func PostChatMessage(
	ctx context.Context,
	l *slog.Logger,
	nowFunc func() time.Time,
	authBusinessLogic auth_businesslogic.ExposedBusinessLogic,
	authUserBusinessLogic authuser_businesslogic.ExposedBusinessLogic,
	adminBusinessLogic admin_businesslogic.ExposedBusinessLogic,
	postedBy string,
	postedByType entity.ChatMessagePostedByType,
	photoStudioID entity.PhotoStudioID,
	userID entity.UserID,
	text string,
	sendPushMessage bool,
) (*entity.ChatMessageWrapper, error) {
	message := entity.ChatMessage{
		Type:         entity.ChatMessageTypeText,
		Text:         text,
		PostedBy:     postedBy,
		PostedByType: postedByType,
		PostedAt:     entity.WTime(nowFunc()),
	}
	if _, err := adminBusinessLogic.CreatePhotoStudioUserChatRoomIFNotExists(
		ctx,
		photoStudioID,
		userID,
	); err != nil {
		return nil, terrors.Wrap(err)
	}
	chatMessage, err := adminBusinessLogic.CreateChatMessage(
		ctx,
		photoStudioID,
		userID,
		&message,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	chatMessageWrappers, err := entity.BuildChatMessageWrapper(
		ctx,
		[]*entity.ChatMessage{
			chatMessage,
		},
		authUserBusinessLogic.GetUsers,
		authBusinessLogic.GetPhotoStudioMembers,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if sendPushMessage {
		notification := entity.Notification{
			ID:                 entity.NotificationID(fmt.Sprintf("%s-%d", entity.NotificationTypeChatMessage, nowFunc().Unix())),
			Type:               entity.NotificationTypeChatMessage,
			ChatMessageWrapper: chatMessageWrappers[0],
		}
		if err := authUserBusinessLogic.PushNotification(ctx, l, userID, &notification); err != nil {
			l.Warn("", "err", err)
		}
		if err := authBusinessLogic.PushNotificationToAllMembers(ctx, l, photoStudioID, &notification); err != nil {
			l.Warn("", "err", err)
		}
	}
	return chatMessageWrappers[0], nil
}
