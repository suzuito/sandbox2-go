package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) APIPostPhotoStudios(ctx *gin.Context) {
	message := struct {
		ID   entity.PhotoStudioID `json:"id"`
		Name string               `json:"name"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		t.P.JSON(ctx, http.StatusBadRequest, ResponseError{})
		return
	}
	dto, err := t.U.APIPostPhotoStudios(ctx, message.ID, message.Name)
	if err != nil {
		var duplicateEntryError *repository.DuplicateEntryError
		if errors.As(err, &duplicateEntryError) {
			t.P.JSON(ctx, http.StatusBadRequest, ResponseError{
				Message: "already exists",
			})
			return
		}
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, ResponseError{})
		return
	}
	t.P.JSON(ctx, http.StatusCreated, dto.PhotoStudio)
}
