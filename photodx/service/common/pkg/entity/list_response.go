package entity

type ListResponse[T any] struct {
	Results    []T  `json:"results"`
	HasNext    bool `json:"hasNext"`
	HasPrev    bool `json:"hasPrev"`
	NextOffset int  `json:"nextOffset"`
	PrevOffset int  `json:"prevOffset"`
}

type ListResponse2[T any] struct {
	Results    []T  `json:"results"`
	HasNext    bool `json:"hasNext"`
	NextOffset int  `json:"nextOffset"`
}
