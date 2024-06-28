package service

import (
	"crypto/md5"
	"fmt"
)

func generatePasswordHash(
	salt []byte,
	password string,
) string {
	enc := md5.New()
	fmt.Fprintf(enc, "%s:%s", string(salt), password)
	return fmt.Sprintf("%x", enc.Sum(nil))
}
