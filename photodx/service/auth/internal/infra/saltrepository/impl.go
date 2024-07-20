package saltrepository

import "context"

type Impl struct{}

func (t *Impl) Get(ctx context.Context) ([]byte, error) {
	return []byte("hoge"), nil
}
