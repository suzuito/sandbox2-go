package usecase

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type XMLURLSet struct {
	XMLName           xml.Name `xml:"urlset"`
	URLs              []XMLURL `xml:"url"`
	XMLNSXsi          string   `xml:"xmlns:xsi,attr"`
	XMLNS             string   `xml:"xmlns,attr"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`
}

type XMLURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
	Lastmod string   `xml:"lastmod"`
}

func (u *UsecaseImpl) GenerateSitemap(
	ctx context.Context,
	siteOrigin string,
	urls *XMLURLSet,
) error {
	limit := 7
	urls.XMLNSXsi = "http://www.w3.org/2001/XMLSchema-instance"
	urls.XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
	urls.XsiSchemaLocation = "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"
	urls.URLs = append(urls.URLs, XMLURL{
		Lastmod: "2023-02-20",
		Loc:     fmt.Sprintf("%s/", siteOrigin),
	})
	offset := 0
	for i := 0; ; i++ {
		primaryKeys := []entity.ArticlePrimaryKey{}
		hasNext := true
		if err := u.RepositoryArticle.SearchArticles(
			ctx,
			SearchArticlesQuery{
				Offset:    offset,
				Limit:     limit,
				SortField: SearchArticlesQuerySortFieldDate,
				SortOrder: SortOrderDesc,
			},
			&primaryKeys,
			&hasNext,
		); err != nil {
			return err
		}
		if len(primaryKeys) <= 0 {
			break
		}
		articles := []entity.Article{}
		if err := u.RepositoryArticle.GetArticlesByPrimaryKey(
			ctx,
			primaryKeys,
			SearchArticlesQuerySortFieldVersion,
			SortOrderDesc,
			&articles,
		); err != nil {
			return err
		}
		for _, article := range articles {
			urls.URLs = append(
				urls.URLs,
				XMLURL{
					Lastmod: article.UpdatedAt.Format("2006-01-02"),
					Loc: fmt.Sprintf(
						"%s/articles/%s",
						siteOrigin,
						article.ID,
					),
				},
			)
		}
		if !hasNext {
			break
		}
		offset = limit * (i + 1)
	}
	return nil
}
