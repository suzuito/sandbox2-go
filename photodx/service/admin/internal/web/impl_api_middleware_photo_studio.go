package web

/*
func (t *Impl) APIMiddlewarePhotoStudio(
	ctx *gin.Context,
) {
	photoStudioID := ctx.Param("photoStudioID")
	dto, err := t.U.APIMiddlewarePhotoStudio(
		ctx,
		entity.PhotoStudioID(photoStudioID),
	)
	if err != nil {
		var noEntryError *repository.NoEntryError
		if errors.As(err, &noEntryError) {
			t.P.JSON(ctx, http.StatusNotFound, common_web.ResponseError{
				Message: fmt.Sprintf("PhotoStudioID '%s' is not found", photoStudioID),
			})
		} else {
			t.L.Error("", "err", err)
			t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
				Message: "internal server error",
			})
		}
		ctx.Abort()
		return
	}
	common_web.CtxSet(ctx, common_web.CtxPhotoStudio, dto.PhotoStudio)
	ctx.Next()
}
*/
