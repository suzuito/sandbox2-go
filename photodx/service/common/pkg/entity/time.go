package entity

import (
	"fmt"
	"time"
)

type WTime time.Time

func (t WTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	return []byte(fmt.Sprintf("%d", tt.UnixMilli())), nil
}
