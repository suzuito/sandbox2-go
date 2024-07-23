package entity

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type ChatMessageWrapper struct {
	ChatMessage
	PostedByStruct PostedByStruct `json:"postedByStruct"`
}

func NewChatMessageWrapper(
	message *ChatMessage,
	user *User,
	member *PhotoStudioMember,
) *ChatMessageWrapper {
	postedByStruct := PostedByStruct{
		Name:            "Unknown",
		ProfileImageURL: "https://vos.line-scdn.net/chdev-console-static/default-profile.png",
	}
	if user != nil {
		postedByStruct.UserID = user.ID
		postedByStruct.Name = user.Name
		postedByStruct.ProfileImageURL = user.ProfileImageURL
	}
	if member != nil {
		postedByStruct.PhotoStudioMemberID = member.ID
		postedByStruct.Name = member.Name
		postedByStruct.ProfileImageURL = "https://vos.line-scdn.net/chdev-console-static/default-profile.png"
	}
	return &ChatMessageWrapper{
		ChatMessage:    *message,
		PostedByStruct: postedByStruct,
	}
}

func BuildChatMessageWrapper(
	ctx context.Context,
	chatMessages []*ChatMessage,
	getUsersFunc func(context.Context, []UserID) ([]*User, error),
	getPhotoStudioMembersFunc func(context.Context, []PhotoStudioMemberID) ([]*PhotoStudioMemberWrapper, error),
) ([]*ChatMessageWrapper, error) {
	usersMap := map[UserID][]*User{}
	{
		userIDs := arrayutil.Map(
			arrayutil.Filter(
				chatMessages,
				func(chatMessage *ChatMessage) bool {
					return chatMessage.PostedByType == ChatMessagePostedByTypeUser
				},
			),
			func(chatMessage *ChatMessage) UserID {
				return UserID(chatMessage.PostedBy)
			},
		)
		userIDs = arrayutil.Uniq(userIDs)
		if len(userIDs) > 0 {
			users, err := getUsersFunc(ctx, userIDs)
			if err != nil {
				return nil, terrors.Wrap(err)
			}
			usersMap = arrayutil.ListToMap(users, func(u *User) UserID {
				return u.ID
			})
		}
	}
	photoStudioMembersMap := map[PhotoStudioMemberID][]*PhotoStudioMemberWrapper{}
	{
		photoStudioMemberIDs := arrayutil.Map(
			arrayutil.Filter(
				chatMessages,
				func(chatMessage *ChatMessage) bool {
					return chatMessage.PostedByType == ChatMessagePostedByTypePhotoStudioMember
				},
			),
			func(chatMessage *ChatMessage) PhotoStudioMemberID {
				return PhotoStudioMemberID(chatMessage.PostedBy)
			},
		)
		photoStudioMemberIDs = arrayutil.Uniq(photoStudioMemberIDs)
		if len(photoStudioMemberIDs) > 0 {
			photoStudioMembers, err := getPhotoStudioMembersFunc(ctx, photoStudioMemberIDs)
			if err != nil {
				return nil, terrors.Wrap(err)
			}
			photoStudioMembersMap = arrayutil.ListToMap(
				photoStudioMembers,
				func(m *PhotoStudioMemberWrapper) PhotoStudioMemberID {
					return m.ID
				},
			)
		}
	}
	wmessages := arrayutil.Map(
		chatMessages,
		func(m *ChatMessage) *ChatMessageWrapper {
			var user *User
			var photoStudioMember *PhotoStudioMember
			switch m.PostedByType {
			case ChatMessagePostedByTypeUser:
				users := usersMap[UserID(m.PostedBy)]
				if len(users) > 0 {
					user = users[0]
				}
			case ChatMessagePostedByTypePhotoStudioMember:
				photoStudioMembers := photoStudioMembersMap[PhotoStudioMemberID(m.PostedBy)]
				if len(photoStudioMembers) > 0 {
					photoStudioMember = photoStudioMembers[0].PhotoStudioMember
				}
			default:
			}
			return NewChatMessageWrapper(
				m,
				user,
				photoStudioMember,
			)
		},
	)
	return wmessages, nil
}
