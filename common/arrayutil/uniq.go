package arrayutil

func Uniq[T comparable](coll []T) []T {
	returned := []T{}
	for _, elem := range coll {
		isDup := false
		for _, v := range returned {
			if v == elem {
				isDup = true
				break
			}
		}
		if isDup {
			continue
		}
		returned = append(returned, elem)
	}
	return returned
}
