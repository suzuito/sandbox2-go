package usecase

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOPostAdminFiles struct {
}

func (t *Impl) PostAdminFiles(
	ctx context.Context,
	fileName string,
	fileType entity.FileType,
	input io.Reader,
) (*DTOPostAdminFiles, error) {
	fileUploaded, err := t.S.CreateFileUploaded(ctx, fileName, fileType, input)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if _, err := t.S.StartFileUploadedProcess(ctx, fileUploaded); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPostAdminFiles{}, nil
}
