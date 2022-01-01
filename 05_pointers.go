package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)

	i := 42
	p = &i

	fmt.Println(p)
}
