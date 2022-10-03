package main

import (
	"fmt"
)

type Value interface {
	int64 | float64
}

type Key interface {
	string | int64
}

func main() {
	intMap := map[string]int64 {
		"first": 87,
		"second": 99,
	}
	floatMap := map[string]float64 {
		"first": 88.34,
		"second": 34.45,
	}
	floatMapWithIntKey := map[int64]float64 {
		1: 88.34,
		2: 34.45,
	}

	fmt.Printf("Non-generic Sums: %v, and %v\n",SumInts(intMap),SumFloats(floatMap))
	fmt.Printf("Generic Sums: %v, and %v\n",SumIntsOrFloats[string,int64](intMap),SumIntsOrFloats[string, float64](floatMap))
	fmt.Printf("Generic Sums(omit type param): %v, and %v\n",SumIntsOrFloats(intMap),SumIntsOrFloats(floatMap))
	fmt.Printf("Generic Sums(constraint with interface): %v, and %v\n",SumIntsOrFloatsWithInterface(intMap),SumIntsOrFloatsWithInterface(floatMapWithIntKey))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _,v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V{
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloatsWithInterface[K Key, V Value](m map[K]V) V{
	var s V
	for _, v := range m {
		s += v
	}
	return s
}