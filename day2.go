package main

import(
	"fmt"
	"os"
	"io"
	"bufio"
	"log"
	"strings"
)

func pt1(r io.Reader) {

	validPasswords := 0

	for {
		// valid if the count of the letter is in the range [min, max]
		// 1-3 a: abcde
		var min    int
		var max    int
		var letter string
		var pwd    string

		_, err := fmt.Fscanf(r, "%d-%d %s %s\n", &min, &max, &letter, &pwd)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		occurrences := strings.Count(pwd, string(letter[0]))
		if occurrences >= min && occurrences <= max { validPasswords += 1 }

	}

	fmt.Printf("pt1 valid passwords: %d\n", validPasswords)
}

func pt2(r io.Reader) {

	validPasswords := 0

	for {
		// valid if only 1 of these 1-indexed positions has the letter
		// 1-3 a: abcde
		var idx1    int
		var idx2    int
		var letter string
		var pwd    string

		_, err := fmt.Fscanf(r, "%d-%d %s %s\n", &idx1, &idx2, &letter, &pwd)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		idx1_yes := pwd[idx1 - 1] == letter[0]
		idx2_yes := pwd[idx2 - 1] == letter[0]
		if idx1_yes && !idx2_yes || !idx1_yes && idx2_yes { validPasswords += 1 }

	}

	fmt.Printf("pt2 valid passwords: %d\n", validPasswords)
}

func main() {
	filename  := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file);

	pt1(r)
	file.Seek(0, 0)
	pt2(r)

	file.Close()
}