package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelPhotoStudioMember struct {
	ID            entity.PhotoStudioMemberID `gorm:"primaryKey;not null"`
	PhotoStudioID entity.PhotoStudioID       `gorm:"not null"`
	Email         string                     `gorm:"not null"`
	Name          string                     `gorm:"not null"`
	Active        bool                       `gorm:"not null;default:false"`
	CreatedAt     time.Time                  `gorm:"not null"`
	UpdatedAt     time.Time                  `gorm:"not null"`

	// Associationsの種類によってforeignKeyの意味が変わるからややこしい。belongsToとhasMany
	// (belongsTo)PhotoStudioMemberはPhotoStudioに属する。
	// `gorm:"foreignKey:PhotoStudioID"`のPhotoStudioIDは、modelPhotoStudioMember(このテーブル)のPhotoStudioIDを指す。
	PhotoStudio *modelPhotoStudio `gorm:"foreignKey:PhotoStudioID"`
	// (hasMany)PhotoStudioMemberは複数のRoleを持つ。
	// `gorm:"foreignKey:PhotoStudioMemberID"`のPhotoStudioMemberIDはRole(先のテーブル)のPhotoStudioMemberIDを指す。
	Roles modelPhotoStudioMemberRoles `gorm:"foreignKey:PhotoStudioMemberID"`
}

func (t *modelPhotoStudioMember) TableName() string {
	return "photo_studio_members"
}

func (t *modelPhotoStudioMember) ToEntity() *entity.PhotoStudioMember {
	return &entity.PhotoStudioMember{
		ID:            t.ID,
		PhotoStudioID: t.PhotoStudioID,
		Email:         t.Email,
		Name:          t.Name,
		Active:        t.Active,
	}
}

func newModelPhotoStudioMember(
	s *entity.PhotoStudioMember,
) *modelPhotoStudioMember {
	return &modelPhotoStudioMember{
		ID:            s.ID,
		PhotoStudioID: s.PhotoStudioID,
		Email:         s.Email,
		Name:          s.Name,
		Active:        s.Active,
	}
}
