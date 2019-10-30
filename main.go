package main

import (
	"fmt"
)

var (
	m = []string{"+", "-", "*"}
	n = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func main() {
	fmt.Print(g(1, 1, 1, 1))
}

func g(f, l, o, r int) string {
	m := map[int]string{
		0: fmt.Sprintf("%v %s %v", l, m[o], n[r]),
		1: fmt.Sprintf("%v %s %v", n[l], m[o], r),
	}
	return m[f]
}
