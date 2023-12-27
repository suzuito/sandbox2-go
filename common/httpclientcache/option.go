package httpclientcache

type ClientOption struct {
	KeyGen    *KeyGen
	TTLInDays int
}
