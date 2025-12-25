package main

func MakeVcMoveMap() map[byte]string {
	m := make(map[byte]string)

	m['j'] = "U"
	m['f'] = "U'"

	m['s'] = "D"
	m['l'] = "D'"

	m['h'] = "F"
	m['g'] = "F'"

	m['w'] = "B"
	m['o'] = "B'"

	m['i'] = "R"
	m['k'] = "R'"

	m['d'] = "L"
	m['e'] = "L'"

	return m
}
