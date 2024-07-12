package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) GetLineLink(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
) (*entity.LineLinkInfo, error) {
	info, err := t.Repository.GetLineLinkInfo(ctx, photoStudioID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return info, nil
}

func (t *Impl) ActivateLineLink(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
) (*entity.LineLinkInfo, error) {
	info := entity.LineLinkInfo{
		PhotoStudioID: photoStudioID,
	}
	created, err := t.Repository.CreateLineLinkInfo(ctx, &info)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}

func (t *Impl) DeactivateLineLink(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
) error {
	if err := t.Repository.DeleteLineLinkInfo(ctx, photoStudioID); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) SetLineLinkInfoMessagingAPIChannelSecret(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	secret string,
) (*entity.LineLinkInfo, error) {
	updated, err := t.Repository.SetLineLinkInfoMessagingAPIChannelSecret(ctx, photoStudioID, secret)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return updated, nil
}
