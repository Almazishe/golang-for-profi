package main

import "fmt"

func Print[T any](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	ints := []int{1, 2, 3, 4}
	strings := []string{"123", "345", "dsvf"}
	Print(ints)
	Print(strings)
}
