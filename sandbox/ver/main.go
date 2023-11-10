package main

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

func main() {
	v1 := version.Must(version.NewVersion("v1.2.3-alpha.0"))
	fmt.Println(
		v1.Segments(),
		v1.Prerelease(),
	)
}
