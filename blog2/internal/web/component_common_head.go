package web

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

const SiteName = "otiuzu pages"

type SiteMetaData struct {
	OGP       OGPData
	Canonical string
	LDJSON    []LDJSONData
}

type OGPData struct {
	Title       string
	Description string
	Locale      string
	Type        string
	URL         string
	SiteName    string
	Image       string
}

type LDJSONData struct {
	Context          string            `json:"@context"`
	Type             string            `json:"@type"`
	Headline         string            `json:"headline,omitempty"`
	Description      string            `json:"description,omitempty"`
	MainEntityOfPage string            `json:"mainEntityOfPage,omitempty"`
	DatePublished    string            `json:"datePublished,omitempty"`
	Author           *LDJSONDataAuthor `json:"author,omitempty"`
	ItemListElement  []LDJSONItem      `json:"itemListElement,omitempty"`
}

type LDJSONItem struct {
	Type     string `json:"@type"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	Item     string `json:"item,omitempty"`
}

type LDJSONDataAuthor struct {
	Type string `json:"@type"`
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

func NewSiteMetaDataFromContext(
	ctx *gin.Context,
	siteOrigin string,
	title string,
	description string,
	t string,
	image string,
) *SiteMetaData {
	return &SiteMetaData{
		OGP: OGPData{
			Title:       title,
			Description: description,
			Locale:      "ja_JP",
			Type:        t,
			URL:         NewPageURLFromContext(ctx, siteOrigin),
			Image:       image,
			SiteName:    SiteName,
		},
		Canonical: NewPageURLFromContext(ctx, siteOrigin),
	}
}

func NewPageURLFromContext(ctx *gin.Context, siteOrigin string) string {
	u, _ := url.Parse(siteOrigin)
	if ctx.Request != nil {
		u.Path = ctx.Request.URL.Path
	}
	return u.String()
}

func NewPageURL(siteOrigin string, path string) string {
	u, _ := url.Parse(siteOrigin)
	u.Path = path
	return u.String()
}

type ComponentCommonHead struct {
	Title              string
	Meta               *SiteMetaData
	GoogleTagManagerID string
}
