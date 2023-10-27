package constant

import "time"

var JST *time.Location

func init() {
	JST, _ = time.LoadLocation("Asia/Tokyo")
}
