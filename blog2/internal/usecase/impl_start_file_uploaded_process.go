package usecase

import (
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) StartFileUploadedProcessFromGCF(ctx context.Context, data []byte) error {
	file := entity.FileUploaded{}
	if err := json.Unmarshal(data, &file); err != nil {
		return terrors.Wrap(err)
	}
	_, err := t.S.StartFileUploadedProcess(ctx, &file)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
