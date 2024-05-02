package web

type Breadcrumb struct {
	Path   string
	URL    string
	Name   string
	NoLink bool
}

type Breadcrumbs []Breadcrumb

func (t Breadcrumbs) LDJSON() LDJSONData {
	itemListElement := []LDJSONItem{}
	for i, v := range t {
		url := v.URL
		if v.NoLink {
			url = ""
		}
		itemListElement = append(
			itemListElement,
			LDJSONItem{
				Type:     "ListItem",
				Position: i + 1,
				Name:     v.Name,
				Item:     url,
			},
		)
	}
	return LDJSONData{
		Context:         "https://schema.org",
		Type:            "BreadcrumbList",
		ItemListElement: itemListElement,
	}
}
