package web

type ResponseSearch[T any] struct {
	Results []T `json:"results"`
}
