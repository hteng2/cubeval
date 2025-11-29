package main

import (
	"fmt"
)

func main() {
	cube := NewCubeDefault()

	fmt.Println(cube.printable(White))

	fmt.Println(cube.printable(Yellow))

	fmt.Println(cube.printable(Green))

	fmt.Println(cube.printable(Blue))

	fmt.Println(cube.printable(Red))

	fmt.Println(cube.printable(Orange))

	fmt.Println(cube.printable(Black))

	fmt.Println(cube.printable(Grey))

	fmt.Println("Cubeval explorer")
}
