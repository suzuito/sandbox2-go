package entity

import "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"

type PhotoStudioMemberID string

type PhotoStudioMember struct {
	ID            PhotoStudioMemberID `json:"id"`
	PhotoStudioID PhotoStudioID       `json:"photoStudioId"`
	Email         string              `json:"email"`
	Name          string              `json:"name"`
	Active        bool                `json:"active"`
}

func (t *PhotoStudioMember) Validate() error {
	// TODO email validation
	// TODO name validation
	return nil
}

const MaxRolesPerPhotoStudioMember = 10

type PhotoStudioMemberWrapper struct {
	*PhotoStudioMember
	Roles       []*rbac.Role
	PhotoStudio *PhotoStudio
}
