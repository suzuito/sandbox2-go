package arrayutil

func Filter[Elem any](arr []Elem, cond func(e Elem) bool) []Elem {
	ret := []Elem{}
	for i := range arr {
		if cond(arr[i]) {
			ret = append(ret, arr[i])
		}
	}
	return ret
}
