package businesslogic

import (
	"context"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type BusinessLogic interface {
	// impl_photo_studio.go
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*common_entity.PhotoStudio, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		name string,
	) (*common_entity.PhotoStudio, error)

	// impl_photo_studio_member.go
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		email string,
		name string,
	) (*common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, string, error)
	GetPhotoStudioMember(
		ctx context.Context,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
	) (*common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, error)
	VerifyPhotoStudioMemberPassword(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		email string,
		password string,
	) (*common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, error)

	// impl_admin_access_token.go
	CreateAdminAccessToken(
		ctx context.Context,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
	) (string, error)

	// impl_admin_refresh_token.go
	CreateAdminRefreshToken(
		ctx context.Context,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
	) (string, error)
	VerifyAdminRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.AdminPrincipalRefreshToken, error)

	// impl_web_push.go
	GetWebPushVAPIDPublicKey(
		ctx context.Context,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
	) (string, error)
	CreateWebPushSubscription(
		ctx context.Context,
		subscription *webpush.Subscription,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
	) error
}
