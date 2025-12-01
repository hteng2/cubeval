package main

import (
	"fmt"
)

func main() {
	cube := NewCubeDefault()

	fmt.Println(cube.stringSide())

	fmt.Println(cube.stringTop())

	fmt.Println("Cubeval explorer")
}
