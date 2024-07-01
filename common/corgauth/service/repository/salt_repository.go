package repository

import "context"

type SaltRepository interface {
	// Get an salt from repository
	// If salt does not exists return EntityNotFoundError.
	Get(ctx context.Context) ([]byte, error)
}
