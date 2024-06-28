package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) APIPostPhotoStudioMembers(ctx *gin.Context) {
	message := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		t.P.JSON(ctx, http.StatusBadRequest, ResponseError{
			Message: "parameter error",
		})
		return
	}
	photoStudio := ctxGet[*entity.PhotoStudio](ctx, ctxPhotoStudio)
	dto, err := t.U.APIPostPhotoStudioMembers(ctx, photoStudio.ID, message.Email, message.Name)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, ResponseError{})
		return
	}
	t.P.JSON(ctx, http.StatusCreated, struct {
		SentInvitation bool `json:"sentInvitation"`
	}{
		SentInvitation: dto.SentInvitation,
	})
}
