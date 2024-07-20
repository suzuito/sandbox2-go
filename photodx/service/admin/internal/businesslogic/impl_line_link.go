package businesslogic

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func (t *Impl) GetActiveLineLink(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
) (*entity.LineLinkInfo, error) {
	info, err := t.Repository.GetLineLinkInfo(ctx, photoStudioID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if !info.Active {
		return nil, &common_repository.NoEntryError{
			EntryType: "LineLinkInfo",
			EntryID:   string(photoStudioID),
		}
	}
	return info, nil
}

func (t *Impl) ActivateLineLink(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
) (*entity.LineLinkInfo, error) {
	info := entity.LineLinkInfo{
		PhotoStudioID: photoStudioID,
		Active:        true,
	}
	returnedLineLinkInfo, err := t.Repository.CreateLineLinkInfo(ctx, &info)
	if err != nil {
		var noDuplicateEntryError *common_repository.DuplicateEntryError
		if !errors.As(err, &noDuplicateEntryError) {
			return nil, terrors.Wrap(err)
		}
		returnedLineLinkInfo, err = t.Repository.SetLineLinkInfoActive(ctx, photoStudioID, true)
		if err != nil {
			return nil, terrors.Wrap(err)
		}
	}
	return returnedLineLinkInfo, nil
}

func (t *Impl) DeactivateLineLink(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
) (*entity.LineLinkInfo, error) {
	info, err := t.Repository.SetLineLinkInfoActive(ctx, photoStudioID, false)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return info, nil
}

func (t *Impl) SetLineLinkInfo(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	arg *repository.SetLineLinkInfoArgument,
) (*entity.LineLinkInfo, error) {
	updated, err := t.Repository.SetLineLinkInfo(ctx, photoStudioID, arg)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return updated, nil
}
