package entity

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func BuildChatRoomWrappers(
	ctx context.Context,
	chatRooms []*ChatRoom,
	getUsersFunc func(context.Context, []UserID) ([]*User, error),
	getPhotoStudiosFunc func(context.Context, []PhotoStudioID) ([]*PhotoStudio, error),
) ([]*ChatRoomWrapper, error) {
	usersMap := map[UserID][]*User{}
	{
		userIDs := arrayutil.MapUniq(chatRooms, func(r *ChatRoom) UserID { return r.UserID })
		users, err := getUsersFunc(ctx, userIDs)
		if err != nil {
			return nil, terrors.Wrap(err)
		}
		usersMap = arrayutil.ListToMap(users, func(u *User) UserID { return u.ID })
	}
	photoStudiosMap := map[PhotoStudioID][]*PhotoStudio{}
	{
		photoStudioIDs := arrayutil.MapUniq(chatRooms, func(r *ChatRoom) PhotoStudioID { return r.PhotoStudioID })
		photoStudios, err := getPhotoStudiosFunc(ctx, photoStudioIDs)
		if err != nil {
			return nil, terrors.Wrap(err)
		}
		photoStudiosMap = arrayutil.ListToMap(photoStudios, func(u *PhotoStudio) PhotoStudioID { return u.ID })
	}
	chatRoomWrappers := arrayutil.Map(
		chatRooms,
		func(chatRoom *ChatRoom) *ChatRoomWrapper {
			var photoStudio *PhotoStudio
			{
				photoStudios, exists := photoStudiosMap[chatRoom.PhotoStudioID]
				if exists {
					photoStudio = photoStudios[0]
				}
			}
			var user *User
			{
				users, exists := usersMap[chatRoom.UserID]
				if exists {
					user = users[0]
				}
			}
			return &ChatRoomWrapper{
				ChatRoom:    chatRoom,
				PhotoStudio: photoStudio,
				User:        user,
			}
		},
	)
	return chatRoomWrappers, nil
}
