package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPageURL(t *testing.T) {
	setting := &WebSetting{
		SiteOrigin: "https://example.com",
	}

	router := gin.Default()
	router.GET("/page", func(c *gin.Context) {
		pageURL := getPageURL(c, setting)
		c.String(http.StatusOK, pageURL)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/page", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "https://example.com/page", w.Body.String())
}

func TestNewSiteMetaData(t *testing.T) {
	setting := &WebSetting{
		SiteOrigin: "https://example.com",
	}

	router := gin.Default()
	router.GET("/metadata", func(c *gin.Context) {
		metaData := newSiteMetaData(c, setting, "Test Page", "This is a test page", "website", "https://example.com/image.jpg")
		c.JSON(http.StatusOK, metaData)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/metadata", nil)
	router.ServeHTTP(w, req)

	expectedResponse := siteMetaData{
		OGP: ogpData{
			Title:       "Test Page",
			Description: "This is a test page",
			Locale:      "ja_JP",
			Type:        "website",
			URL:         "https://example.com/metadata",
			Image:       "https://example.com/image.jpg",
			SiteName:    siteName,
		},
		Canonical: "https://example.com/metadata",
	}

	var actualResponse siteMetaData
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
