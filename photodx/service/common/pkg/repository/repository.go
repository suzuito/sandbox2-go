package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
)

type Repository interface {
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
	) (*entity.PhotoStudio, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudio *entity.PhotoStudio,
	) (*entity.PhotoStudio, error)

	GetPhotoStudioMember(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	GetPhotoStudioMemberByEmail(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		photoStudioMember *entity.PhotoStudioMember,
		initialPasswordHashValue string,
		initialRoles []rbac.RoleID,
	) (*entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)
	GetPhotoStudioMemberPasswordHashByEmail(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
	) (string, *entity.PhotoStudioMember, []*rbac.Role, *entity.PhotoStudio, error)

	CreateOAuth2State(
		ctx context.Context,
		state *oauth2loginflow.State,
	) (*oauth2loginflow.State, error)
	GetAndDeleteOAuth2State(
		ctx context.Context,
		stateCode oauth2loginflow.StateCode,
	) (*oauth2loginflow.State, error)

	CreateUser(
		ctx context.Context,
		user *entity.User,
	) (*entity.User, error)
	GetUserByResourceOwnerID(
		ctx context.Context,
		providerID oauth2loginflow.ProviderID,
		resourceOwnerID oauth2loginflow.ResourceOwnerID,
	) (*entity.User, error)
}
