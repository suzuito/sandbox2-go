package arrayutil

func ListToMap[Key comparable, Elem any](arr []Elem, conv func(e Elem) Key) map[Key][]Elem {
	ret := map[Key][]Elem{}
	for i := range arr {
		k := conv(arr[i])
		ret[k] = append(ret[k], arr[i])
	}
	return ret
}
