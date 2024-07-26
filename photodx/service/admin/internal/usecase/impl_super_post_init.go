package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIPostSuperInit struct {
	PhotoStudio     *common_entity.PhotoStudio       `json:"photoStudio"`
	SuperMember     *common_entity.PhotoStudioMember `json:"superMember"`
	InitialPassword string                           `json:"initialPassword"`
}

func (t *Impl) APIPostSuperInit(ctx context.Context) (*DTOAPIPostSuperInit, error) {
	photoStudio, err := t.AuthBusinessLogic.CreatePhotoStudio(
		ctx,
		"godzilla",
		"XXX Photo Studio",
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	superMember, _, _, initialPassword, err := t.AuthBusinessLogic.CreatePhotoStudioMember(
		ctx,
		photoStudio.ID,
		"super@photodx.tach.dev",
		"スーパーユーザー",
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostSuperInit{
		PhotoStudio:     photoStudio,
		SuperMember:     superMember,
		InitialPassword: initialPassword,
	}, nil
}
