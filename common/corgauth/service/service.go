package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
)

type Service interface {
	// impl_organization.go
	CreateOrganization(
		ctx context.Context,
		organizationID entity.OrganizationID,
	) (*entity.Organization, error)

	// impl_principal.go
	CreatePrincipal(
		ctx context.Context,
		organizationID entity.OrganizationID,
		email string,
	) (*entity.Principal, string, error)
	VerifyPrincipalPassword(
		ctx context.Context,
		organizationID entity.OrganizationID,
		email string,
		password string,
	) (*entity.Principal, []*entity.Role, error)

	// impl_refresh_token.go
	CreateRefreshToken(
		ctx context.Context,
		principalID entity.PrincipalID,
	) (string, error)
	VerifyRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.ClaimsRefreshToken, error)

	// impl_access_token.go
	CreateAccessToken(
		ctx context.Context,
		principalID entity.PrincipalID,
	) (string, error)
	VerifyAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.ClaimsAccessToken, error)
}
