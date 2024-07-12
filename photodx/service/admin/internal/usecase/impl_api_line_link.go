package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIGetLINELink struct {
	LineLinkInfo *entity.LineLinkInfo `json:"lineLinkInfo"`
}

func (t *Impl) APIGetLINELink(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
) (*DTOAPIGetLINELink, error) {
	info, err := t.BusinessLogic.GetLineLink(ctx, principal.GetPhotoStudioID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetLINELink{
		LineLinkInfo: info,
	}, nil
}

type DTOAPIPostLINELink struct {
	LineLinkInfo *entity.LineLinkInfo `json:"lineLinkInfo"`
}

func (t *Impl) APIPostLINELink(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
) (*DTOAPIPostLINELink, error) {
	info, err := t.BusinessLogic.ActivateLineLink(ctx, principal.GetPhotoStudioID())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostLINELink{
		LineLinkInfo: info,
	}, nil
}

func (t *Impl) APIDeleteLINELink(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
) error {
	if err := t.BusinessLogic.DeactivateLineLink(ctx, principal.GetPhotoStudioID()); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

type DTOAPIPutLINELinkMessagingAPIChannelSecret struct {
	LineLinkInfo *entity.LineLinkInfo `json:"lineLinkInfo"`
}

func (t *Impl) APIPutLINELinkMessagingAPIChannelSecret(
	ctx context.Context,
	principal common_entity.AdminPrincipalAccessToken,
	secret string,
) (*DTOAPIPutLINELinkMessagingAPIChannelSecret, error) {
	info, err := t.BusinessLogic.SetLineLinkInfoMessagingAPIChannelSecret(ctx, principal.GetPhotoStudioID(), secret)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPutLINELinkMessagingAPIChannelSecret{
		LineLinkInfo: info,
	}, nil
}
