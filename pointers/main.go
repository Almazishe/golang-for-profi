package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}


func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointer(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	
	fP := &f
	fmt.Println("Memory addres of f:", fP)
	fmt.Println("Value of f:", *fP)

	processPointer(fP)
	fmt.Println("Memory addres of f:", fP)
	fmt.Println("Value of f:", *fP)


	x := returnPointer(f)
	fmt.Println("Memory addres of x:", x)
	fmt.Println("Value of x:", *x)
	fmt.Println("Memory address of f:", &f)
	fmt.Println("Value of f:", f)


	xx := bothPointer(fP)
	fmt.Println("Value of xx:", *xx)
	fmt.Println("Memory addres of xx:", xx)
}