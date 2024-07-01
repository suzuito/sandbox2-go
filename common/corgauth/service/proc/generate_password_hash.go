package proc

type GeneratePasswordHashFunc func(
	plainPassword string,
	salt []byte,
) string
