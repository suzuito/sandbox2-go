package arrayutil

func Map[InputT any, OutputT any](input []InputT, converter func(InputT) OutputT) []OutputT {
	returned := []OutputT{}
	for _, v := range input {
		returned = append(returned, converter(v))
	}
	return returned
}
