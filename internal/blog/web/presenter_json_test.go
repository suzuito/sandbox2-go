package web

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPresenterJSONResponse(t *testing.T) {
	p := PresenterJSON{}
	w := httptest.NewRecorder()
	r := gin.New()
	ctx := gin.CreateTestContextOnly(w, r)
	p.Response(ctx, &PresenterArgJSON{
		Code: 200,
		Obj:  gin.H{"hoge": "fuga"},
	})
}
