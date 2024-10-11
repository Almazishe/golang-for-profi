package main

import "fmt"

func main() {
	//var a []int
	//
	//for i := 0; i < 2049; i++ {
	//	a = append(a, i)
	//	fmt.Println("idx", i, "cap:", cap(a), "len:", len(a))
	//}

	t := make([]int, 4)
	t = append(t, 12)
	fmt.Println(t)
}
