package arrayutil

import (
	"slices"
)

func Uniq[E comparable](src []E) []E {
	ret := []E{}
	for _, a := range src {
		if !slices.Contains(ret, a) {
			ret = append(ret, a)
		}
	}
	return ret
}
