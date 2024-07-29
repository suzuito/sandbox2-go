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
