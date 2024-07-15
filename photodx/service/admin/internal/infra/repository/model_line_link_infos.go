package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelLineLinkInfo struct {
	PhotoStudioID             common_entity.PhotoStudioID `gorm:"primaryKey;not null"`
	MessagingAPIChannelSecret string                      `gorm:"column:messaging_api_channel_secret;not null"`
	Active                    bool                        `gorm:"not null"`
	CreatedAt                 time.Time                   `gorm:"not null"`
	UpdatedAt                 time.Time                   `gorm:"not null"`
}

func (t *modelLineLinkInfo) TableName() string {
	return "line_link_infos"
}

func (t *modelLineLinkInfo) ToEntity() *entity.LineLinkInfo {
	return &entity.LineLinkInfo{
		PhotoStudioID:             t.PhotoStudioID,
		MessagingAPIChannelSecret: t.MessagingAPIChannelSecret,
		Active:                    t.Active,
	}
}

func newModelLineLinkInfo(s *entity.LineLinkInfo) *modelLineLinkInfo {
	return &modelLineLinkInfo{
		PhotoStudioID:             s.PhotoStudioID,
		MessagingAPIChannelSecret: s.MessagingAPIChannelSecret,
		Active:                    s.Active,
	}
}
