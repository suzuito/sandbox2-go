package web

import (
	"testing"
)

func TestPresenterRedirectResponse(t *testing.T) {
	// なぜかだめ
	// p := PresenterRedirect{}
	// w := httptest.NewRecorder()
	// r := gin.New()
	// ctx := gin.CreateTestContextOnly(w, r)
	// p.Response(ctx, &PresenterArgRedirect{
	// 	Code:     http.StatusFound,
	// 	Location: "https://www.example.com/admin",
	// })
}
