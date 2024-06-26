package auth0

import "context"

type Validator interface {
	ValidateToken(ctx context.Context, tokenString string) (any, error)
}
