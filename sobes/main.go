package main

import "fmt"

type Parent struct{}

func (c *Parent) Print() {
	fmt.Println("parent")
}

type Child struct {
	Parent
}

func (p *Child) Print() {
	fmt.Println("child")
}

func main() {
	var x Child

	{
		var x = 951

		a := string(x)

		fmt.Println(a)
	}

	x.Print()
	x.Parent.Print()

	var a float64 = -1.5
	fmt.Println(a / 0)
}
