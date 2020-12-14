package main

import(
	"fmt"
	"os"
	"io"
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

func checkPath(terrain []string, width int, height int, step_width int, step_height int) {
	treeCount := 0

	// x is vertical axis
	// y is horizontal axis
	x := 0
	y := 0
	for x < height {
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
	r := getFileReader(os.Args[1])

	terrain, width, height := readTerrain(r)
	checkPath(terrain, width, height, 1, 1)
	checkPath(terrain, width, height, 3, 1)
	checkPath(terrain, width, height, 5, 1)
	checkPath(terrain, width, height, 7, 1)
	checkPath(terrain, width, height, 1, 2)

}