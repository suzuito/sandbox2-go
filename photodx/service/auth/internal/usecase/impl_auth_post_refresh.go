package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAuthPostRefresh struct {
	AccessToken string
}

func (t *Impl) AuthPostRefresh(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (*DTOAuthPostRefresh, error) {
	accessTokenString, err := t.BusinessLogic.CreateAdminAccessToken(ctx, photoStudioMemberID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAuthPostRefresh{
		AccessToken: accessTokenString,
	}, nil
}
