package main

import(
	"fmt"
	"os"
	"io"
	"bufio"
	"log"
)

func printTerrain(terrain []string) {
	for line := 0; line < len(terrain); line += 1 {
		for _, char := range(terrain[line]) {
			fmt.Printf("%c", char)
		}
		fmt.Println("")
	}
}

func readTerrain(r io.Reader) ([]string, int, int) {
	var terrain []string

	for {
		var line string

		_, err := fmt.Fscanf(r, "%s\n", &line)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		terrain = append(terrain, line[:len(line)])

	}
	// terrain, width, height
	return terrain, len(terrain[0]), len(terrain)
}

func pt1(terrain []string, width int, height int, step_width int, step_height int) {
	// printTerrain(terrain)
	// fmt.Printf("width: %d, height: %d\n", width, height)
	treeCount := 0

	// x is vertical axis
	// y is horizontal axis
	x := 0
	y := 0
	for x < height {

		// fmt.Printf("x: %d, y: %d, c: %c\n", x, y, terrain[x][y])
		if terrain[x][y] == '#' {
			treeCount += 1
		}

		x += step_height
		y += step_width
		y = y % width
	}

	fmt.Printf("pt1 treeCount: %d\n", treeCount)
}

func main() {
	filename  := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file);

	terrain, width, height := readTerrain(r)
	pt1(terrain, width, height, 1, 1)
	pt1(terrain, width, height, 3, 1)
	pt1(terrain, width, height, 5, 1)
	pt1(terrain, width, height, 7, 1)
	pt1(terrain, width, height, 1, 2)
	file.Seek(0, 0)
	// pt2(r)

	file.Close()
}