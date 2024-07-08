package repository

import "context"

type SaltRepository interface {
	Get(ctx context.Context) ([]byte, error)
}
