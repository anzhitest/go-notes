package main

import "fmt"

func main() {
	var (
		x int
		y float32
		z string
		p *int
		q bool
	)

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%x\n", p)
	fmt.Printf("%v\n", q)
}
