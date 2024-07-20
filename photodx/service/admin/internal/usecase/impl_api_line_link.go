package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIGetLINELink struct {
	LineLinkInfo *entity.LineLinkInfo `json:"lineLinkInfo"`
}

func (t *Impl) APIGetLINELink(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
) (*DTOAPIGetLINELink, error) {
	info, err := t.BusinessLogic.GetActiveLineLink(ctx, principal.GetPhotoStudioID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetLINELink{
		LineLinkInfo: info,
	}, nil
}

type DTOAPIPutLINELinkActivate struct {
	LineLinkInfo *entity.LineLinkInfo `json:"lineLinkInfo"`
}

func (t *Impl) APIPutLINELinkActivate(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
) (*DTOAPIPutLINELinkActivate, error) {
	info, err := t.BusinessLogic.ActivateLineLink(ctx, principal.GetPhotoStudioID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPutLINELinkActivate{
		LineLinkInfo: info,
	}, nil
}

type DTOAPIPutLINELinkDeactivate struct {
	LineLinkInfo *entity.LineLinkInfo `json:"lineLinkInfo"`
}

func (t *Impl) APIPutLINELinkDeactivate(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
) (*DTOAPIPutLINELinkDeactivate, error) {
	info, err := t.BusinessLogic.DeactivateLineLink(ctx, principal.GetPhotoStudioID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPutLINELinkDeactivate{
		LineLinkInfo: info,
	}, nil
}

type DTOAPIPutLINELinkMessagingAPIChannelSecret struct {
	LineLinkInfo *entity.LineLinkInfo `json:"lineLinkInfo"`
}

func (t *Impl) APIPutLINELink(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
	arg *repository.SetLineLinkInfoArgument,
) (*DTOAPIPutLINELinkMessagingAPIChannelSecret, error) {
	info, err := t.BusinessLogic.SetLineLinkInfo(ctx, principal.GetPhotoStudioID(), arg)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPutLINELinkMessagingAPIChannelSecret{
		LineLinkInfo: info,
	}, nil
}
