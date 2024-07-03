package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIGetInit struct {
	PhotoStudio       *entity.PhotoStudio
	PhotoStudioMember *entity.PhotoStudioMember
}

func (t *Impl) APIGetInit(
	ctx context.Context,
	principal entity.Principal,
) (*DTOAPIGetInit, error) {
	photoStudioMember, _, photoStudio, err := t.B.GetPhotoStudioMember(
		ctx,
		principal.GetPhotoStudioMemberID(),
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetInit{
		PhotoStudio:       photoStudio,
		PhotoStudioMember: photoStudioMember,
	}, nil
}
