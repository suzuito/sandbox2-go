package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
)

// 管理画面へのログイン機能
// 本機能は、ユニットテストなしで良い

const CookieNameAdminSession = "ADMIN_SESSION_ID"

var sessionIDStored = ""

func (t *ControllerImpl) MiddlewareCheckAdminAuth(ctx *gin.Context) {
	argHTML404 := PresenterArgHTML{
		Code: http.StatusNotFound,
		Name: "pc_error.html",
		Obj: responseError{
			Title:   "404 - エンティティが見つかりません",
			Message: "Not found",
		},
	}
	sessionID, err := ctx.Cookie(CookieNameAdminSession)
	if err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{ArgHTML: &argHTML404})
		return
	}
	if sessionID == "" {
		t.Presenters.Response(ctx, PresenterArgStandardError{ArgHTML: &argHTML404})
		return
	}
	if sessionIDStored != sessionID {
		t.Presenters.Response(ctx, PresenterArgStandardError{ArgHTML: &argHTML404})
		return
	}
	ctx.Next()
}

func (t *ControllerImpl) GetAdminLogin(ctx *gin.Context) {
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_login.html",
		Obj:  gin.H{},
	})
}

func (t *ControllerImpl) PostAdminLogin(ctx *gin.Context) {
	password := ctx.PostForm("password")
	if password == "" {
		t.Presenters.Response(ctx, PresenterArgHTML{
			Code: http.StatusOK,
			Name: "admin_login.html",
			Obj:  gin.H{},
		})
		return
	}
	if password != t.Setting.AdminPassword {
		clog.L.Errorf(ctx, "invalid password")
		t.Presenters.Response(ctx, PresenterArgHTML{
			Code: http.StatusOK,
			Name: "admin_login.html",
			Obj:  gin.H{},
		})
		return
	}
	sessionIDStored = password
	ctx.SetCookie(CookieNameAdminSession, sessionIDStored, 86400, "", "", false, false)
	t.Presenters.Response(ctx, PresenterArgRedirect{
		Code:     http.StatusFound,
		Location: "/admin",
	})
}
