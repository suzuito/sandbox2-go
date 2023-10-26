package web

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/common/cweb"
)

type SetRouterFunc func(e *gin.Engine, ctrl *ControllerImpl)

func SetRouter(
	e *gin.Engine,
	ctrl *ControllerImpl,
) {
	setRouterRoot(e, ctrl)
	setRouterApp(e, ctrl)
	setRouterAdmin(e, ctrl)
}

func setRouterRoot(
	e *gin.Engine,
	ctrl *ControllerImpl,
) {
	if ctrl.NoIndex() {
		e.Use(cweb.MiddlewareXRobotsTag())
	}
	e.LoadHTMLGlob(path.Join(ctrl.DirPathTemplates(), "*"))
	e.Static("css", ctrl.DirPathCSS())
	e.Static("images", ctrl.DirPathImages())
	e.NoRoute(ctrl.NoRoute)
	e.GET("robots.txt", ctrl.GetRobots)
	e.GET("sitemap.xml", ctrl.GetSitemap)
	e.GET("health", func(ctx *gin.Context) {
		ctrl.Presenters.Response(ctx, PresenterArgJSON{
			Code: http.StatusOK,
			Obj: gin.H{
				"message": "ok",
			},
		})
	})
	e.GET("test-log", func(ctx *gin.Context) {
		clog.L.Infof(ctx, "hoge")
		clog.L.Errorf(ctx, "fuga")
		clog.L.Debugf(ctx, "foo")
		ctrl.Presenters.Response(ctx, PresenterArgJSON{
			Code: http.StatusOK,
			Obj: gin.H{
				"message": "ok",
			},
		})
	})
}

func setRouterApp(
	e *gin.Engine,
	ctrl *ControllerImpl,
) {
	e.GET("", ctrl.GetTop)
	e.GET("about", ctrl.GetAbout)
	{
		articles := e.Group("articles")
		articles.GET("", ctrl.GetArticles)
		{
			article := articles.Group(":articleID")
			article.Use(ctrl.MiddlewareGetLatestArticle)
			article.GET("", ctrl.GetArticle)
		}
	}
}

func setRouterAdmin(
	e *gin.Engine,
	ctrl *ControllerImpl,
) {
	{
		admin := e.Group("admin")
		{
			adminLogin := admin.Group("login")
			adminLogin.GET("", ctrl.GetAdminLogin)
			adminLogin.POST("", ctrl.PostAdminLogin)
		}
		admin.Use(ctrl.MiddlewareCheckAdminAuth)
		admin.GET("", ctrl.GetAdmin)
		admin.GET("search/source", ctrl.GetAdminSearchSource)
		admin.GET("execute/import-all-sources", ctrl.PostAdminExecuteImportSources)
		{
			adminImport := admin.Group("import/:articleSourceID/:articleSourceVersion")
			adminImport.GET("", ctrl.GetAdminImport)
			adminImport.POST("", ctrl.PostAdminImport)
			adminImport.GET("success", ctrl.GetAdminImportSuccess)
			adminImport.GET("error", ctrl.GetAdminImportError)
		}
		{
			adminArticlesByID := admin.Group("articles/:articleID")
			adminArticlesByID.Use(ctrl.MiddlewareGetArticle)
			adminArticlesByID.GET("", ctrl.GetAdminArticlesByID)
		}
		{
			adminSource := admin.Group("sources/:articleSourceID")
			adminSource.GET("", ctrl.GetAdminSourcesByID)
		}
	}
}
