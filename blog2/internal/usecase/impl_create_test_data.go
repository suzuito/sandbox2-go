package usecase

import (
	"context"
)

func (t *Impl) CreateTestData001(ctx context.Context) error {
	return t.S.CreateTestData(ctx)
}
