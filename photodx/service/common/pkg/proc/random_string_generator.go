package proc

type RandomStringGenerator interface {
	Gen() (string, error)
}
