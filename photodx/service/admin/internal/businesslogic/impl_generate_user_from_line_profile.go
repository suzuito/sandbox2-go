package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) GenerateUserFromLINEProfile(
	ctx context.Context,
	lineLinkInfo *entity.LineLinkInfo,
	lineUserID string,
) (*common_entity.User, error) {
	user, err := t.LINEMessagingAPIClient.GetProfile(
		ctx,
		lineLinkInfo.LongAccessToken,
		lineUserID,
	)
	if err != nil {
		// lineLinkInfo.LongAccessToken is input by user
		// This token is maybe mistoken
		user = &common_entity.User{
			Name:            "名無しさん",
			ProfileImageURL: "",
		}
	}
	user.Active = true
	user.InitializedByUser = false
	return user, nil
}
