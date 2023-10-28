package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog/entity"
	"github.com/suzuito/sandbox2-go/blog/usecase"
)

func (t *ControllerImpl) GetAdmin(ctx *gin.Context) {
	limit := 1000
	query := struct {
		Page int    `form:"page"`
		Tags string `form:"tags"`
	}{}
	if err := ctx.BindQuery(&query); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	tags := strings.Split(query.Tags, ",")
	tagsFiltered := []string{}
	for _, tag := range tags {
		if tag == "" {
			continue
		}
		tagsFiltered = append(tagsFiltered, tag)
	}
	articles := []entity.Article{}
	offset := 0
	if query.Page > 0 {
		offset = query.Page * limit
	}
	hasNext := false
	if err := t.UC.SearchArticles(
		ctx,
		usecase.SearchArticlesQuery{
			Offset:    offset,
			Limit:     limit,
			Tags:      tagsFiltered,
			SortField: usecase.SearchArticlesQuerySortFieldDate,
			SortOrder: usecase.SortOrderDesc,
		},
		&articles,
		&hasNext,
	); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	branches, err := t.RepositoryArticleSource.GetBranches(ctx)
	if err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_top.html",
		Obj: gin.H{
			"Branches": branches,
			"Articles": articles,
		},
	})
}
