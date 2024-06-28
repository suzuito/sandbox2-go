package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

type DTOSuperPostInit struct {
	PhotoStudio     *entity.PhotoStudio
	SuperMember     *entity.PhotoStudioMember
	InitialPassword string
}

func (t *Impl) SuperPostInit(ctx context.Context) (*DTOSuperPostInit, error) {
	photoStudio, err := t.S.CreatePhotoStudio(
		ctx,
		"godzilla",
		"サービス管理",
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	superMember, initialPassword, err := t.S.CreatePhotoStudioMember(
		ctx,
		photoStudio.ID,
		"super@photodx.tach.dev",
		"スーパーユーザー",
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOSuperPostInit{
		PhotoStudio:     photoStudio,
		SuperMember:     superMember,
		InitialPassword: initialPassword,
	}, nil
}
