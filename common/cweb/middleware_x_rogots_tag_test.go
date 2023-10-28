package cweb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareXRobotsTag(t *testing.T) {
	router := gin.New()
	router.Use(MiddlewareXRobotsTag())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, "noindex", w.HeaderMap.Get("X-Robots-Tag"))
}
