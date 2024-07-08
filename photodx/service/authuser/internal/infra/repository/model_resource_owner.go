package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelProviderResourceOwnersUsersMapping struct {
	ProviderID      oauth2loginflow.ProviderID      `gorm:"not null"`
	ResourceOwnerID oauth2loginflow.ResourceOwnerID `gorm:"not null"`
	UserID          entity.UserID                   `gorm:"not null"`
	CreatedAt       time.Time                       `gorm:"not null"`
}

func (t *modelProviderResourceOwnersUsersMapping) TableName() string {
	return `provider_resource_owners_users_mappings`
}
