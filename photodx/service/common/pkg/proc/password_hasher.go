package proc

import (
	"crypto/md5"
	"fmt"
)

type PasswordHasher interface {
	Gen(salt []byte, plainPassword string) string
}

type PasswordHasherMD5 struct{}

func (t *PasswordHasherMD5) Gen(salt []byte, plainPassword string) string {
	enc := md5.New()
	fmt.Fprintf(enc, "%s:%s", string(salt), plainPassword)
	return fmt.Sprintf("%x", enc.Sum(nil))
}
