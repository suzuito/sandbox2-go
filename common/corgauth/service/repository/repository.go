package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
)

type Repository interface {
	// Get an organization from repository
	// If the organization does not exists return EntityNotFoundError.
	GetOrganization(
		ctx context.Context,
		organizationID entity.OrganizationID,
	) (*entity.Organization, error)

	// Create an organization to repository
	// If the organization already exists return EntityAlreadyExistsError.
	CreateOrganization(
		ctx context.Context,
		organization *entity.Organization,
	) (*entity.Organization, error)

	// Create an principal to repository
	// If the principal already exists return EntityAlreadyExistsError.
	GetPrincipal(
		ctx context.Context,
		principalID entity.PrincipalID,
	) (*entity.Principal, []*entity.Role, error)

	// Get an principal from repository
	// If the principal does not exists return EntityNotFoundError.
	GetPrincipalByEmail(
		ctx context.Context,
		organizationID entity.OrganizationID,
		email string,
	) (*entity.Principal, []*entity.Role, error)

	// Create a principal to repository
	// If the principal already exists return EntityAlreadyExistsError.
	CreatePrincipal(
		ctx context.Context,
		organizationID entity.OrganizationID,
		principal *entity.Principal,
		initialPasswordHash string,
		initialRoles []entity.Role,
	) (*entity.Principal, []*entity.Role, error)

	// Get an principal's password hash from repository
	// If the password hash does not exists return EntityNotFoundError.
	GetPrincipalPasswordHash(
		ctx context.Context,
		principalID entity.PrincipalID,
	) (string, error)
}
