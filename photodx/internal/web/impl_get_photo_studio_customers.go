package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) GetPhotoStudioCustomers(ctx *gin.Context) {
	t.P.JSON(ctx, http.StatusOK, ResponseSearch[entity.Customer]{
		Results: []entity.Customer{},
	})
}
