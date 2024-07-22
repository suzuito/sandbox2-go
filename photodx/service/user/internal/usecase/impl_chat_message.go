package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type InputAPIPostPhotoStudioMessages struct {
	Text string `json:"text"`
}

func (t *Impl) APIPostPhotoStudioMessages(
	ctx context.Context,
	principal common_entity.UserPrincipalAccessToken,
	photoStudioID common_entity.PhotoStudioID,
	input *InputAPIPostPhotoStudioMessages,
) (*common_entity.ChatMessageWrapper, error) {
	msg := common_entity.ChatMessage{
		Type:         common_entity.ChatMessageTypeText,
		Text:         input.Text,
		PostedBy:     string(principal.GetUserID()),
		PostedByType: common_entity.ChatMessagePostedByTypeUser,
		PostedAt:     common_entity.WTime(t.NowFunc()),
	}
	created, err := t.AdminBusinessLogic.CreateChatMessage(
		ctx,
		photoStudioID,
		&msg,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.AuthUserBusinessLogic.PushNotification(ctx, t.L, principal.GetUserID(), created.Text); err != nil {
		t.L.Warn("", "err", err)
	}
	a, err := t.attachPostedBy(ctx, []*common_entity.ChatMessage{created})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return a[0], nil
}

type DTOAPIGetPhotoStudioMessages struct {
	Results    []*common_entity.ChatMessageWrapper `json:"results"`
	HasNext    bool                                `json:"hasNext"`
	HasPrev    bool                                `json:"hasPrev"`
	NextOffset int                                 `json:"nextOffset"`
	PrevOffset int                                 `json:"prevOffset"`
}

func (t *Impl) APIGetPhotoStudioMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	listQuery *cgorm.ListQuery,
) (*DTOAPIGetPhotoStudioMessages, error) {
	listQuery.Limit = 30
	if listQuery.Offset < 0 {
		listQuery.Offset = 0
	}
	listQuery.SortColumns = []cgorm.SortColumn{
		{Name: "posted_at", Type: cgorm.Asc},
	}
	messages, hasNext, err := t.AdminBusinessLogic.GetChatMessages(ctx, photoStudioID, listQuery)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	wmessages, err := t.attachPostedBy(ctx, messages)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetPhotoStudioMessages{
		Results:    wmessages,
		HasNext:    hasNext,
		HasPrev:    listQuery.HasPrev(),
		NextOffset: listQuery.NextOffset(),
		PrevOffset: listQuery.PrevOffset(),
	}, nil
}

func (t *Impl) attachPostedBy(
	ctx context.Context,
	messages []*common_entity.ChatMessage,
) ([]*common_entity.ChatMessageWrapper, error) {
	usersMap := map[common_entity.UserID][]*common_entity.User{}
	{
		userIDs := arrayutil.Map(
			arrayutil.Filter(
				messages,
				func(message *common_entity.ChatMessage) bool {
					return message.PostedByType == common_entity.ChatMessagePostedByTypeUser
				},
			),
			func(message *common_entity.ChatMessage) common_entity.UserID {
				return common_entity.UserID(message.PostedBy)
			},
		)
		userIDs = arrayutil.Uniq(userIDs)
		if len(userIDs) > 0 {
			users, err := t.AuthUserBusinessLogic.GetUsers(ctx, userIDs)
			if err != nil {
				return nil, terrors.Wrap(err)
			}
			usersMap = arrayutil.ListToMap(users, func(u *common_entity.User) common_entity.UserID {
				return u.ID
			})
		}
	}
	photoStudioMembersMap := map[common_entity.PhotoStudioMemberID][]*common_entity.PhotoStudioMemberWrapper{}
	{
		photoStudioMemberIDs := arrayutil.Map(
			arrayutil.Filter(
				messages,
				func(message *common_entity.ChatMessage) bool {
					return message.PostedByType == common_entity.ChatMessagePostedByTypePhotoStudioMember
				},
			),
			func(message *common_entity.ChatMessage) common_entity.PhotoStudioMemberID {
				return common_entity.PhotoStudioMemberID(message.PostedBy)
			},
		)
		photoStudioMemberIDs = arrayutil.Uniq(photoStudioMemberIDs)
		if len(photoStudioMemberIDs) > 0 {
			photoStudioMembers, err := t.AuthBusinessLogic.GetPhotoStudioMembers(ctx, photoStudioMemberIDs)
			if err != nil {
				return nil, terrors.Wrap(err)
			}
			photoStudioMembersMap = arrayutil.ListToMap(photoStudioMembers, func(m *common_entity.PhotoStudioMemberWrapper) common_entity.PhotoStudioMemberID {
				return m.ID
			})
		}
	}
	wmessages := arrayutil.Map(
		messages,
		func(m *common_entity.ChatMessage) *common_entity.ChatMessageWrapper {
			var user *common_entity.User
			var photoStudioMember *common_entity.PhotoStudioMember
			switch m.PostedByType {
			case common_entity.ChatMessagePostedByTypeUser:
				users := usersMap[common_entity.UserID(m.PostedBy)]
				if len(users) > 0 {
					user = users[0]
				}
			case common_entity.ChatMessagePostedByTypePhotoStudioMember:
				photoStudioMembers := photoStudioMembersMap[common_entity.PhotoStudioMemberID(m.PostedBy)]
				if len(photoStudioMembers) > 0 {
					photoStudioMember = photoStudioMembers[0].PhotoStudioMember
				}
			default:
			}
			return common_entity.NewChatMessageWrapper(
				m,
				user,
				photoStudioMember,
			)
		},
	)
	return wmessages, nil
}
