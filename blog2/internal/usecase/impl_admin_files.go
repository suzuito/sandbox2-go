package usecase

type DTOPostAdminFiles struct {
}

// func (t *Impl) PostAdminFiles(
// 	ctx context.Context,
// 	fileName string,
// 	fileType entity.FileType,
// 	input io.Reader,
// ) (*DTOPostAdminFiles, error) {
// 	fileUploaded, err := t.S.CreateFileUploaded(ctx, fileName, fileType, input)
// 	if err != nil {
// 		return nil, terrors.Wrap(err)
// 	}
// 	if _, err := t.S.StartFileUploadedProcess(ctx, fileUploaded); err != nil {
// 		return nil, terrors.Wrap(err)
// 	}
// 	return &DTOPostAdminFiles{}, nil
// }
