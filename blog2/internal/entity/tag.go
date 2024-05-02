package entity

type TagID string

func NewTagIDs(tagIDs []string) []TagID {
	ret := make([]TagID, len(tagIDs))
	for i := range tagIDs {
		ret[i] = TagID(tagIDs[i])
	}
	return ret
}

type Tag struct {
	ID   TagID  `json:"id"`
	Name string `json:"name"`
}
