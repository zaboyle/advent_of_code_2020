package main

import(
	"fmt"
	"io"
	"os"
	"log"
)

func getSeatLocations(r io.Reader) []string {

	var seatLocations []string

	for {
		var seat string

		_, err := fmt.Fscanf(r, "%s\n", &seat)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		seatLocations = append(seatLocations, seat[:len(seat)])
	}
	return seatLocations
}

func getSeatID(seatLocation string) uint32 {
	var row uint32 = 0
	var col uint32 = 0

	var lbRow uint32 = 0
	var ubRow uint32 = 127

	for _, char := range seatLocation[:7] {
		// F for Front
		var mid uint32 = (ubRow + lbRow) / 2
		if char == 'F' {
			ubRow = mid
			row = ubRow
		} else if char == 'B' {
			lbRow = mid + 1
			row = lbRow
		}
	}

	var lbCol uint32 = 0
	var ubCol uint32 = 7

	for _, char := range seatLocation[7:] {
		// L for Left
		var mid uint32 = (ubCol + lbCol) / 2
		if char == 'L' {
			ubCol = mid
			col = ubCol
		} else if char == 'R' {
			lbCol = mid + 1
			col = lbCol
		}
	}

	// seat ID formula
	return row * 8 + col
}

func min(seatIDs []uint32) uint32 {
	// max uint32
	var min uint32 = ^uint32(0)
	for _, i := range seatIDs {
		if i < min {
			min = i
		}
	}
	return min
}

func max(seatIDs []uint32) uint32 {
	var max uint32 = 0
	for _, i := range seatIDs {
		if i > max {
			max = i
		}
	}
	return max
}

func getSeatIDs(seatLocations []string) []uint32 {
	var seatIDs []uint32
	for _, seat := range seatLocations {
		seatID := getSeatID(seat)
		seatIDs = append(seatIDs, seatID)
	}
	return seatIDs
}

func pt1(seatIDs []uint32) {
	max := max(seatIDs)
	fmt.Printf("highest seat ID: %d\n", max)
}

func pt2(seatIDs []uint32) {
	// find missing in list
	min := min(seatIDs)
	max := max(seatIDs)

	set := make(map[uint32]bool)

	for i := min; i < max; i++ {
		set[i] = false
	}

	for _, val := range seatIDs {
		set[val] = true
	}

	for j := min; j < max; j++ {
		if !set[j] {
			fmt.Printf("missing seat ID: %d\n", j)
		}
	}

}

func main() {
	r             := getFileReader(os.Args[1])
	seatLocations := getSeatLocations(r)
	seatIDs       := getSeatIDs(seatLocations)
	pt1(seatIDs)
	pt2(seatIDs)
}