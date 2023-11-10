package main

import "fmt"

func main() {
	main1()
}

func main5() {
	FuncA(1, int64(1))
	FuncA("hoge", float64(1))
	FuncA(struct{}{}, float64(1))
	// FuncA([]int64{1, 2, 3}, float64(1))
	// FuncA(map[string]string{"hoge": "fuga"}, float64(1))
}

func FuncA[K comparable, V int64 | float64](a1 K, a2 V) {}

func main4() {
	f1(int64(1))
	f2(int64(1))
	// f1(X(1))
	f2(X(1))
}

func f1[N int64](n N)  {}
func f2[N ~int64](n N) {}

type X int64

func main3() {
	var c ComputerInterface[string, int64]
	c = &Computer[string, int64]{
		m: map[string]int64{},
	}
	c.Set("hoge", 1)
	c.Set("fuga", 2)
	fmt.Println(c.Sum())
}

type A[V comparable] int64
type B A[string]

type ComputerInterface[K comparable, V int64 | float64] interface {
	Set(k K, v V)
	Sum() V
}

type Computer[K comparable, V int64 | float64] struct {
	m map[K]V
}

func (t *Computer[K, V]) Set(k K, v V) {
	t.m[k] = v
}

func (t *Computer[K, V]) Sum() V {
	var s V
	for _, v := range t.m {
		s += v
	}
	return s
}

type Sum[N int | float32] struct {
	Arr []N
}

func (t *Sum[N]) Total() N {
	var returned N = 0
	for _, v := range t.Arr {
		returned += v
	}
	return N(returned)
}

func main2() {
	fmt.Println(
		AddNumber[int](1, 2),
		AddNumber[float32](float32(1.0), float32(2.0)),
	)
	fmt.Println(
		AddNumber(1, 2),
		AddNumber(float32(1.0), float32(2.0)),
	)
}

func AddNumber[N MyNumber](a, b N) N {
	return a + b
}

type MyNumber interface {
	int | float32
}

func AddNumberInt(a, b int) int {
	return a + b
}

func AddNumberFloat32(a, b float32) float32 {
	return a + b
}

func main1() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Printf(
		"%v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
