package web

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPresenterHTMLResponse(t *testing.T) {
	p := PresenterHTML{}
	w := httptest.NewRecorder()
	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	ctx := gin.CreateTestContextOnly(w, r)
	p.Response(ctx, &PresenterArgHTML{
		Code: 1,
		Name: "n1",
		Obj:  gin.H{"hoge": "fuga"},
	})
}
