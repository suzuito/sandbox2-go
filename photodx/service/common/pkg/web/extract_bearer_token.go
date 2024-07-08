package web

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(ctx *gin.Context) string {
	authorizationHeaderString := ctx.GetHeader("Authorization")
	authorizationHeaderParts := strings.Fields(authorizationHeaderString)
	if len(authorizationHeaderParts) != 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
		return ""
	}
	accessToken := authorizationHeaderParts[1]
	return accessToken
}
