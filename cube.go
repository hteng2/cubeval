package main

import (
	"fmt"
)

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

type Slice uint8

const (
	M Slice = iota
	S
	E
)

type Edge struct {
	colors      [2]Color
	orientation uint8
}

type Corner struct {
	colors      [3]Color
	orientation uint8
}

type Center struct {
	color Color
}

type Cube struct {
	edges   [12]Edge
	corners [8]Corner
	centers [6]Center
}

func InitCube() Cube {
	var edges [12]Edge

	var slice Slice

	// M slice
	slice = M
	for UD := range uint8(2) {
		for FB := range uint8(2) {
			index := uint8(slice) + 3*UD + 6*FB

			var edge Edge

			if UD == 0 {
				edge.colors[0] = White
			} else {
				edge.colors[0] = Yellow
			}

			if FB == 0 {
				edge.colors[1] = Green
			} else {
				edge.colors[1] = Blue
			}

			edge.orientation = 0

			edges[index] = edge
		}
	}

	// S slice
	slice = S
	for UD := range uint8(2) {
		for RL := range uint8(2) {
			index := uint8(slice) + 3*UD + 6*RL

			var edge Edge

			if UD == 0 {
				edge.colors[0] = White
			} else {
				edge.colors[0] = Yellow
			}

			if RL == 0 {
				edge.colors[1] = Red
			} else {
				edge.colors[1] = Orange
			}

			edge.orientation = 0

			edges[index] = edge
		}
	}

	// E slice
	slice = E
	for FB := range uint8(2) {
		for RL := range uint8(2) {
			index := uint8(slice) + 3*FB + 6*RL

			var edge Edge

			if FB == 0 {
				edge.colors[0] = Green
			} else {
				edge.colors[0] = Blue
			}

			if RL == 0 {
				edge.colors[1] = Red
			} else {
				edge.colors[1] = Orange
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
				} else {
					corner.colors[0] = Yellow
				}

				if FB == 0 {
					corner.colors[1] = Green
				} else {
					corner.colors[1] = Blue
				}

				if RL == 0 {
					corner.colors[2] = Red
				} else {
					corner.colors[2] = Orange
				}

				if (UD+FB+RL)%2 == 0 {
					tmp := corner.colors[1]
					corner.colors[1] = corner.colors[2]
					corner.colors[2] = tmp
				}

				corner.orientation = 0

				corners[index] = corner
			}
		}
	}

	centers := [6]Center{
		{White},
		{Yellow},
		{Green},
		{Blue},
		{Red},
		{Orange},
	}

	return Cube{edges, corners, centers}
}

func (cube *Cube) cycleEdges(indices [4]uint8, flip bool) {
	tmp := cube.edges[indices[3]]
	cube.edges[indices[3]] = cube.edges[indices[2]]
	cube.edges[indices[2]] = cube.edges[indices[1]]
	cube.edges[indices[1]] = cube.edges[indices[0]]
	cube.edges[indices[0]] = tmp

	if flip {
		for _, i := range indices {
			cube.edges[i].orientation = 1 - cube.edges[i].orientation
		}
	}
}

func (cube *Cube) cycleCorners(indices [4]uint8, twist uint8) {
	tmp := cube.corners[indices[3]]
	cube.corners[indices[3]] = cube.corners[indices[2]]
	cube.corners[indices[2]] = cube.corners[indices[1]]
	cube.corners[indices[1]] = cube.corners[indices[0]]
	cube.corners[indices[0]] = tmp

	cube.corners[indices[0]].orientation += twist
	cube.corners[indices[0]].orientation %= 3

	cube.corners[indices[1]].orientation += 3 - twist
	cube.corners[indices[1]].orientation %= 3

	cube.corners[indices[2]].orientation += twist
	cube.corners[indices[2]].orientation %= 3

	cube.corners[indices[3]].orientation += 3 - twist
	cube.corners[indices[3]].orientation %= 3
}

func (cube *Cube) cycleCenters(indices [4]uint8) {
	tmp := cube.centers[indices[3]]
	cube.centers[indices[3]] = cube.centers[indices[2]]
	cube.centers[indices[2]] = cube.centers[indices[1]]
	cube.centers[indices[1]] = cube.centers[indices[0]]
	cube.centers[indices[0]] = tmp
}

func (cube *Cube) doU() {
	cube.cycleEdges(
		[4]uint8{0, 7, 6, 1},
		false,
	)
	cube.cycleCorners(
		[4]uint8{0, 4, 6, 2},
		0,
	)
}

func (cube *Cube) doD() {
	cube.cycleEdges(
		[4]uint8{3, 4, 9, 10},
		false,
	)
	cube.cycleCorners(
		[4]uint8{0, 4, 6, 2},
		0,
	)
}

func (cube *Cube) doF() {
	cube.cycleEdges(
		[4]uint8{0, 2, 3, 8},
		true,
	)
	cube.cycleCorners(
		[4]uint8{0, 1, 5, 4},
		2,
	)
}

func (cube *Cube) doB() {
	cube.cycleEdges(
		[4]uint8{6, 11, 9, 5},
		true,
	)
	cube.cycleCorners(
		[4]uint8{2, 6, 7, 3},
		1,
	)
}

func (cube *Cube) doR() {
	cube.cycleEdges(
		[4]uint8{1, 5, 4, 2},
		false,
	)
	cube.cycleCorners(
		[4]uint8{0, 2, 3, 1},
		1,
	)
}

func (cube *Cube) doL() {
	cube.cycleEdges(
		[4]uint8{7, 8, 10, 11},
		false,
	)
	cube.cycleCorners(
		[4]uint8{4, 5, 7, 6},
		2,
	)
}

func (cube *Cube) DoMove(move string) error {
	n := 1

	if len(move) == 2 {
		switch move[1] {
		case '2':
			n = 2
		case '\'':
			n = 3
		default:
			return fmt.Errorf("cube.DoMove: illegal move %s", move)
		}
	}

	if len(move) > 2 {
		return fmt.Errorf("cube.DoMove: illegal move %s", move)
	}

	for range n {
		switch move[0] {
		case 'U':
			cube.doU()
		case 'D':
			cube.doD()
		case 'F':
			cube.doF()
		case 'B':
			cube.doB()
		case 'R':
			cube.doR()
		case 'L':
			cube.doL()
		default:
			return fmt.Errorf("cube.DoMove: illegal move %s", move)
		}
	}
	return nil
}

func (cube *Cube) stringSide() string {
	//    -- -- --
	//   -- -- --/
	//  -- -- --//
	// ## ## ##///
	// ## ## ##//
	// ## ## ##/

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

	e := &cube.edges
	c := &cube.corners
	h := &cube.centers

	result := ""

	result += fmt.Sprintf("   %s-- %s-- %s--\r\n",
		colors[c[6].colors[c[6].orientation]],
		colors[e[6].colors[e[6].orientation]],
		colors[c[2].colors[c[2].orientation]])

	result += fmt.Sprintf("  %s-- %s-- %s--%s/\r\n",
		colors[e[7].colors[e[7].orientation]],
		colors[h[0].color],
		colors[e[1].colors[e[1].orientation]],
		colors[c[2].colors[(2+c[2].orientation)%3]])

	result += fmt.Sprintf(" %s-- %s-- %s--%s/%s/\r\n",
		colors[c[4].colors[c[4].orientation]],
		colors[e[0].colors[e[0].orientation]],
		colors[c[0].colors[c[0].orientation]],
		colors[e[1].colors[1-e[1].orientation]],
		colors[e[5].colors[1-e[5].orientation]])

	result += fmt.Sprintf("%s## %s## %s##%s/%s/%s/\r\n",
		colors[c[4].colors[(1+c[4].orientation)%3]],
		colors[e[0].colors[1-e[0].orientation]],
		colors[c[0].colors[(2+c[0].orientation)%3]],
		colors[c[0].colors[(1+c[0].orientation)%3]],
		colors[h[4].color],
		colors[c[3].colors[(1+c[3].orientation)%3]])

	result += fmt.Sprintf("%s## %s## %s##%s/%s/\r\n",
		colors[e[8].colors[e[8].orientation]],
		colors[h[2].color],
		colors[e[2].colors[e[2].orientation]],
		colors[e[2].colors[1-e[2].orientation]],
		colors[e[4].colors[1-e[4].orientation]])

	result += fmt.Sprintf("%s## %s## %s##%s/\r\n",
		colors[c[5].colors[(2+c[5].orientation)%3]],
		colors[e[3].colors[1-e[3].orientation]],
		colors[c[1].colors[(1+c[1].orientation)%3]],
		colors[c[1].colors[(2+c[1].orientation)%3]])

	result += reset
	return result
}

func (cube *Cube) stringTop() string {
	//   -- -- --
	// | ## ## ## |
	// | ## ## ## |
	// | ## ## ## |
	//   -- -- --

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

	e := &cube.edges
	c := &cube.corners
	h := &cube.centers

	result := ""

	result += fmt.Sprintf("  %s-- %s-- %s--\r\n",
		colors[c[6].colors[(2+c[6].orientation)%3]],
		colors[e[6].colors[1-e[6].orientation]],
		colors[c[2].colors[(1+c[2].orientation)%3]])

	result += fmt.Sprintf("%s| %s## %s## %s## %s|\r\n",
		colors[c[6].colors[(1+c[6].orientation)%3]],
		colors[c[6].colors[c[6].orientation]],
		colors[e[6].colors[e[6].orientation]],
		colors[c[2].colors[c[2].orientation]],
		colors[c[2].colors[(2+c[2].orientation)%3]])

	result += fmt.Sprintf("%s| %s## %s## %s## %s|\r\n",
		colors[e[7].colors[1-e[7].orientation]],
		colors[e[7].colors[e[7].orientation]],
		colors[h[0].color],
		colors[e[1].colors[e[1].orientation]],
		colors[e[1].colors[1-e[1].orientation]])

	result += fmt.Sprintf("%s| %s## %s## %s## %s|\r\n",
		colors[c[4].colors[(2+c[4].orientation)%3]],
		colors[c[4].colors[c[4].orientation]],
		colors[e[0].colors[e[0].orientation]],
		colors[c[0].colors[c[0].orientation]],
		colors[c[0].colors[(1+c[0].orientation)%3]])

	result += fmt.Sprintf("  %s-- %s-- %s--\r\n",
		colors[c[4].colors[(1+c[4].orientation)%3]],
		colors[e[0].colors[1-e[0].orientation]],
		colors[c[0].colors[(2+c[0].orientation)%3]])

	result += reset

	return result
}
