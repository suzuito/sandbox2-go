package usecase

import "context"

type Usecase interface {
	MiddlewareAccessTokenAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOMiddlewareAccessTokenAuthe, error)
	GetInit(ctx context.Context) (*DTOGetInit, error)
}
