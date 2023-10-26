package web

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

const siteName = "otiuzu pages"

type siteMetaData struct {
	OGP       ogpData
	Canonical string
	LDJSON    []ldjsonData
}

type ogpData struct {
	Title       string
	Description string
	Locale      string
	Type        string
	URL         string
	SiteName    string
	Image       string
}

type ldjsonData struct {
	Context          string           `json:"@context"`
	Type             string           `json:"@type"`
	Headline         string           `json:"headline,omitempty"`
	Description      string           `json:"description,omitempty"`
	MainEntityOfPage string           `json:"mainEntityOfPage,omitempty"`
	DatePublished    string           `json:"datePublished,omitempty"`
	Author           ldjsonDataAuthor `json:"author,omitempty"`
}

type ldjsonDataAuthor struct {
	Type string `json:"@type"`
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

func newSiteMetaData(
	ctx *gin.Context,
	setting *WebSetting,
	title string,
	description string,
	t string,
	image string,
) *siteMetaData {
	return &siteMetaData{
		OGP: ogpData{
			Title:       title,
			Description: description,
			Locale:      "ja_JP",
			Type:        t,
			URL:         getPageURL(ctx, setting),
			Image:       image,
			SiteName:    siteName,
		},
		Canonical: getPageURL(ctx, setting),
	}
}

func getPageURL(ctx *gin.Context, setting *WebSetting) string {
	u, _ := url.Parse(setting.SiteOrigin)
	if ctx.Request != nil {
		u.Path = ctx.Request.URL.Path
	}
	return u.String()
}
