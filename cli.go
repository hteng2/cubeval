package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)

const (
	sig_normal byte = iota
	sig_exit
)

type Term struct {
	rows int
	cols int

	history []string

	cmdHandler *func(*Term, string) byte

	fd       int
	oldState *term.State
}

func InitTerm(cmdHandler *func(*Term, string) byte) *Term {
	// save cursor position
	fmt.Print("\033[s")
	
	var t Term

	// dimensions
	cols, rows, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	t.rows = rows
	t.cols = cols

	// history
	t.history = []string{}

	// clear terminal and move cursor
	for range rows {
		fmt.Println()
	}

	err = t.MoveCursor(0, 0)
	if err != nil {
		panic(err)
	}

	// command handler
	t.cmdHandler = cmdHandler

	// enter raw mode
	t.fd = int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(t.fd)
	if err != nil {
		panic(err)
	}

	t.oldState = oldState

	return &t
}

func (t *Term) Restore() {
	err := term.Restore(t.fd, t.oldState)
	if err != nil {
		panic(err)
	}
	fmt.Print("\033[u")
}

func (t *Term) MoveCursor(row int, col int) error {
	if t == nil {
		return errors.New("moveCursor: nil term")
	}

	if row < 0 || t.rows <= row || col < 0 || t.cols <= col {
		return errors.New("moveCursor: position out of bounds")
	}

	fmt.Printf("\033[%d;%dH", row, col)

	return nil
}

func (t *Term) Loop() {
	reader := bufio.NewReader(os.Stdin)

	buffer := []byte{}

	fmt.Print("> ")

	for {
		buf := make([]byte, 3)
		n, err := reader.Read(buf)
		if err != nil {
			return
		}
		if n != 1 {
			continue
		}

		b := buf[0]

		switch b {
		case 13: // enter
			fmt.Print("\r\n")

			if string(buffer) == "exit" {
				return
			}

			t.history = append(t.history, string(buffer)+"\r\n")

			signal := (*t.cmdHandler)(t, string(buffer))
			if signal == sig_exit {
				return
			}
			buffer = []byte{}
			fmt.Print("> ")
		case 127: // backspace
			if len(buffer) > 0 {
				fmt.Print("\b \b")
				buffer = buffer[:len(buffer)-1]
			}
		default:
			if 32 <= b && b < 128 { // if printable
				fmt.Printf("%c", b)
				buffer = append(buffer, b)
			}
		}
	}
}