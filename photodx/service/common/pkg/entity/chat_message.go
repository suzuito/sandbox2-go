package entity

type ChatMessageID string
type ChatMessageType string
type ChatMessagePostedByType string
type ChatMessage struct {
	ID           ChatMessageID           `json:"id"`
	Type         ChatMessageType         `json:"type"`
	Text         string                  `json:"text"`
	PostedBy     string                  `json:"postedBy"`
	PostedByType ChatMessagePostedByType `json:"postedByType"`
	PostedAt     WTime                   `json:"postedAt"`
}

const (
	ChatMessageTypeText ChatMessageType = "text"
)

const (
	ChatMessagePostedByTypeUser              ChatMessagePostedByType = "user"
	ChatMessagePostedByTypePhotoStudioMember ChatMessagePostedByType = "photoStudioMember"
)

type PostedByStruct struct {
	PhotoStudioMemberID PhotoStudioMemberID `json:"photoStudioMemberId"`
	UserID              UserID              `json:"userId"`
	Name                string              `json:"name"`
	ProfileImageURL     string              `json:"profileImageUrl"`
}

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
