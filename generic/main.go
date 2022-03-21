package main

import "fmt"

func main() {
	
	ints := map[string]int64{"first": 23, "second": 45}
	floats := map[string]float64{"first": 45.3, "second": 89.23}
	fmt.Printf("Generics sums %v and %v\n", sumIntsOrFloats(ints), sumIntsOrFloats(floats))
}

func sumIntsOrFloats[K string, V int64 | float64](m map[K]V)V {
	var s V;

	for _, v := range m {
		s += v;
	}

	return s;
}