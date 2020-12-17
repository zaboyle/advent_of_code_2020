package main

import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
)

func readInstructions(r io.Reader) []string {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(bytes), "\n")
	return instructions
}

func pt1(instructions []string) {
	pc  := 0
	acc := 0
	instr_visited := make(map[int]bool)
	for i, _ := range instructions {
		instr_visited[i] = false
	}

	// stop once we reach a line we've seen
	for !instr_visited[pc] {
		// acc, nop, jmp
		switch opcode := instructions[pc][:3]; opcode {
		case "acc":
			instr_visited[pc] = true
			amount, err := strconv.Atoi(instructions[pc][4:])
			if err != nil { log.Fatal(err) }
			acc += amount
			pc++

		case "nop":
			instr_visited[pc] = true
			pc++

		case "jmp":
			instr_visited[pc] = true
			amount, err := strconv.Atoi(instructions[pc][4:])
			if err != nil { log.Fatal(err) }
			pc += amount

		default:
			log.Fatal("invalid instruction!!!")
		}

	}

	fmt.Printf("instruction %d hit twice, acc value is %d\n", pc, acc)
}

func pt2(instructions []string) {

}

func main() {
	r := getFileReader(os.Args[1])
	instructions := readInstructions(r)

	pt1(instructions)
	pt2(instructions)
}