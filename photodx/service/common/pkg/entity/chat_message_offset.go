package entity

import "time"

type ChatMessageOffset struct {
	PostedAt time.Time     `form:"postedAt"`
	ID       ChatMessageID `form:"id"`
}
