package main

import "fmt"

func main() {
	a := make([]int, 5, 5)
	a2 := a
	a2 = asd(a2, 1)
	a[2] = 5
	a = asd(a, 2)
	println()
	fmt.Println(len(a), cap(a), a[:cap(a)])
	fmt.Println(len(a2), cap(a2), a2[:cap(a2)])
}

func asd(a []int, i int) []int {
	return append(a, i)
}
