package businesslogic

import (
	"context"
	"net/http"
	"net/url"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
)

type BusinessLogic interface {
	// impl_photo_studio.go
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
	) (*entity.PhotoStudio, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		name string,
	) (*entity.PhotoStudio, error)

	// impl_photo_studio_member.go
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		name string,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, string, error)
	GetPhotoStudioMember(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	VerifyPhotoStudioMemberPassword(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		password string,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)

	// impl_admin_access_token.go
	CreateAdminAccessToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyAdminAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.AdminPrincipal, error)

	// impl_admin_refresh_token.go
	CreateAdminRefreshToken(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (string, error)
	VerifyAdminRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.AdminPrincipalRefreshToken, error)

	// impl_user_access_token.go
	CreateUserAccessToken(
		ctx context.Context,
		userID entity.UserID,
	) (string, error)
	VerifyUserAccessToken(ctx context.Context,
		accessToken string,
	) (entity.UserPrincipal, error)

	// impl_user_refresh_token.go
	CreateUserRefreshToken(
		ctx context.Context,
		userID entity.UserID,
	) (string, error)
	VerifyUserRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.UserPrincipalRefreshToken, error)

	// impl_oauth2_login_flow.go
	CreateOAuth2State(
		ctx context.Context,
		providerID oauth2loginflow.ProviderID,
		callbackURL *url.URL,
		oauth2RedirectURL *url.URL,
	) (*oauth2loginflow.State, error)
	CallbackVerifyState(
		ctx context.Context,
		stateCode oauth2loginflow.StateCode,
	) (*oauth2loginflow.State, error)
	FetchAccessToken(
		ctx context.Context,
		req *http.Request,
	) (string, error)
	FetchProfileAndCreateUserIfNotExists(
		ctx context.Context,
		accessToken string,
		providerID oauth2loginflow.ProviderID,
		fetchProfile oauth2loginflow.Oauth2FetchProfileFunc,
	) (*entity.User, error)
}
