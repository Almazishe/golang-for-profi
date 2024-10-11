package main

import "fmt"

func change(s []string) {
	s[0] = "change_func"
}

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	var S0 = a[0:1]
	fmt.Println("S0", S0)
	S0[0] = "S0"

	fmt.Println("a", a)
	fmt.Println("S0", S0, "len", len(S0), "cap", cap(S0))

	var S12 = a[1:3]
	fmt.Println("S12", S12)
	S12[0] = "S12_0"
	S12[1] = "S12_1"

	fmt.Println("a", a)
	fmt.Println("S12", S12, "len", len(S12), "cap", cap(S12))

	change(S12)

	fmt.Println("a", a)
	fmt.Println("S12", S12, "len", len(S12), "cap", cap(S12))

	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")

	a[0] = "-N1"

	fmt.Println("a", a, "len", len(a), "cap", cap(a))
	fmt.Println("S0", S0, "len", len(S0), "cap", cap(S0))
	fmt.Println("S12", S12, "len", len(S12), "cap", cap(S12))

	S0 = append(S0, "N4")

	fmt.Println("a", a, "len", len(a), "cap", cap(a))
	fmt.Println("S0", S0, "len", len(S0), "cap", cap(S0))
	fmt.Println("S12", S12, "len", len(S12), "cap", cap(S12))

	a[0] = "-N1-"
	a[1] = "-N2-"

	fmt.Println("a", a, "len", len(a), "cap", cap(a))
	fmt.Println("S0", S0, "len", len(S0), "cap", cap(S0))
	fmt.Println("S12", S12, "len", len(S12), "cap", cap(S12))
}
