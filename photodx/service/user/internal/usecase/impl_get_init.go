package usecase

import (
	"context"
	"fmt"
)

type DTOGetInit struct{}

func (t *Impl) GetInit(ctx context.Context) (*DTOGetInit, error) {
	return nil, fmt.Errorf("not impl")
}
