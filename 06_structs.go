package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 10
	fmt.Println(v)

	var (
		v2 = Vertex{X: 15}
		v3 = Vertex{}
	)
	fmt.Println(v2, v3)

}
