package main

import (
	"context"
)

/*
func main() {
	c := gen(1, 2)
	o := sq(c)
	fmt.Println(<-o)
	fmt.Println(<-o)
}

func gen(nums ...int32) <-chan int32 {
	out := make(chan int32)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func sq(in <-chan int32) <-chan int32 {
	out := make(chan int32)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
*/

func main() {
	ctx := context.Background()

	go func() {
		get(ctx)
	}()
	go func() {
		get(ctx)
	}()
}

func get(ctx context.Context) error {
	return nil
}
