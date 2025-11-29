package main

import "fmt"

type Color uint8

const (
	White Color = iota
	Yellow
	Green
	Blue
	Red
	Orange
	Black
	Grey
)

type Face uint8

const (
	U Face = iota
	D
	F
	B
	R
	L
)

type Edge struct {
	colors      [2]Color
	position    [2]Face
	orientation uint8
}

type Corner struct {
	colors      [3]Color
	position    [3]Face
	orientation uint8
}

type Center struct {
	color    Color
	position Face
}

type Cube struct {
	edges   [12]Edge
	corners [8]Corner
	centers [6]Center
}

func NewCubeDefault() Cube {
	var edges [12]Edge

	var slice uint8

	// M slice
	slice = 0
	for UD := range uint8(2) {
		for FB := range uint8(2) {
			index := slice + 3*UD + 6*FB

			var edge Edge

			if UD == 0 {
				edge.colors[0] = White
				edge.position[0] = U
			} else {
				edge.colors[0] = Yellow
				edge.position[0] = D
			}

			if FB == 0 {
				edge.colors[1] = Green
				edge.position[1] = F
			} else {
				edge.colors[1] = Blue
				edge.position[1] = B
			}

			edge.orientation = 0

			edges[index] = edge
		}
	}

	// S slice
	slice = 1
	for UD := range uint8(2) {
		for RL := range uint8(2) {
			index := slice + 3*UD + 6*RL

			var edge Edge

			if UD == 0 {
				edge.colors[0] = White
				edge.position[0] = U
			} else {
				edge.colors[0] = Yellow
				edge.position[0] = D
			}

			if RL == 0 {
				edge.colors[1] = Red
				edge.position[1] = R
			} else {
				edge.colors[1] = Orange
				edge.position[1] = L
			}

			edge.orientation = 0

			edges[index] = edge
		}
	}

	// E slice
	slice = 2
	for FB := range uint8(2) {
		for RL := range uint8(2) {
			index := slice + 3*FB + 6*RL

			var edge Edge

			if FB == 0 {
				edge.colors[0] = Green
				edge.position[0] = F
			} else {
				edge.colors[0] = Blue
				edge.position[0] = B
			}

			if RL == 0 {
				edge.colors[1] = Red
				edge.position[1] = R
			} else {
				edge.colors[1] = Orange
				edge.position[1] = L
			}

			edge.orientation = 0

			edges[index] = edge
		}
	}

	var corners [8]Corner

	for UD := range uint8(2) {
		for FB := range uint8(2) {
			for RL := range uint8(2) {
				index := UD + 2*FB + 4*RL

				var corner Corner

				if UD == 0 {
					corner.colors[0] = White
					corner.position[0] = U
				} else {
					corner.colors[0] = Yellow
					corner.position[0] = D
				}

				if FB == 0 {
					corner.colors[1] = Green
					corner.position[1] = F
				} else {
					corner.colors[1] = Blue
					corner.position[1] = B
				}

				if RL == 0 {
					corner.colors[2] = Red
					corner.position[2] = R
				} else {
					corner.colors[2] = Orange
					corner.position[2] = L
				}

				corner.orientation = 0

				corners[index] = corner
			}
		}
	}

	centers := [6]Center{
		{White, U},
		{Yellow, D},
		{Green, F},
		{Blue, B},
		{Red, R},
		{Orange, L},
	}

	return Cube{edges, corners, centers}
}

func (cube Cube) printable(color Color) string {
	//    .. .. ..
	//   .. .. .. -
	//  .. .. .. --
	// ## ## ## ---
	// ## ## ## --
	// ## ## ## -

	colors := [8]string{
		"\033[1;37m",
		"\033[1;33m",
		"\033[1;32m",
		"\033[1;34m",
		"\033[1;31m",
		"\033[0;33m",
		"\033[0;30m",
		"\033[0;37m",
	}

	reset := "\033[0m"

	return fmt.Sprintln(colors[color], cube, reset)
}
