package saltrepository

import "context"

func (t *Impl) Get(ctx context.Context) ([]byte, error) {
	// TODO 後でSecretManagerを使うように作り変える
	return []byte("hoge"), nil
}
