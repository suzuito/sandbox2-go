package web

import (
	"github.com/gin-gonic/gin"
)

type responseError struct {
	Title   string
	Message string
	Header  struct{}
	Meta    siteMetaData
}

type Presenter interface {
	Response(
		ctx *gin.Context,
		code int,
		name string,
		obj any,
	)
}
