package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HelpCmd() {
	fmt.Print("?    --- displays commands list\r\n")
	fmt.Print("exit --- exit\r\n")
	fmt.Print("do   --- do move\r\n")
	fmt.Print("vc   --- enter virtual cube mode")
	fmt.Print("view --- view cube\r\n")
}

func DoMoveCmd(tokens *[]string, cube *Cube) {
	cubeCopy := *cube
	for i, token := range *tokens {
		if i == 0 {
			continue
		}

		err := cube.DoMove(token)

		if err != nil {
			fmt.Printf("%s\r\n", err.Error())
			*cube = cubeCopy
			return
		}
	}
}

func VCCmd(t *Term, cube *Cube) {
	reader := bufio.NewReader(os.Stdin)

	cubeCopy := *cube
	movemap := MakeVcMoveMap()

	exit := false
	for !exit {
		ViewCmd(cube)

		b, err := reader.ReadByte()

		if err != nil {
			panic(err)
		}

		switch b {
		case 13: // enter
			exit = true
		case 27: // escape
			*cube = cubeCopy
			exit = true
		default:
			move, exists := movemap[b]
			if exists {
				cube.DoMove(move)
			}
		}
		fmt.Print("\033[11A")
	}
	for range 11 {
		fmt.Print("\033[2K\r\n")
	}
	fmt.Print("\033[11A")
}

func ViewCmd(cube *Cube) {
	fmt.Print(cube.stringSide())
	fmt.Print(cube.stringTop())
}

func MakeCmdHandler(cube *Cube) func(t *Term, cmd string) byte {
	return func(t *Term, cmd string) byte {
		tokens := strings.Split(cmd, " ")

		switch tokens[0] {
		case "?":
			HelpCmd()
		case "do":
			DoMoveCmd(&tokens, cube)
		case "vc":
			VCCmd(t, cube)
		case "view":
			ViewCmd(cube)
		default:
			fmt.Printf("unknown command: %s\r\n", tokens[0])
		}
		return sig_normal
	}
}
