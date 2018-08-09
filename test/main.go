package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	x     int
	y     int
	h     bool
	ans   = 0
	flags [][]bool
)

func init() {
	flag.IntVar(&x, "x", 20, "x is rows")
	flag.IntVar(&y, "y", 20, "y is cols")
}

// isGreatThan21 return true if (x, y) more than 21
func isLessThan21(x, y int) bool {
	xstr := strconv.Itoa(x)
	ystr := strconv.Itoa(y)

	sum := 0
	for _, sx := range xstr {
		sum += int(sx - '0')
	}
	for _, sy := range ystr {
		sum += int(sy - '0')
	}

	if sum < 21 {
		return true
	}
	return false
}

// getNumberOfPoint return the number of the points
func getNumberOfPoint(startX, startY int) int {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			tx := startX + dx
			ty := startY + dy

			if tx >= 0 && ty >= 0 && tx < x && ty < y && isLessThan21(tx, ty) && flags[tx][ty] == false {
				flags[tx][ty] = true
				ans++
				getNumberOfPoint(tx, ty)
			}
		}
	}
	return ans
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	flags = make([][]bool, x)
	for i := 0; i < x; i++ {
		subArray := make([]bool, y)
		for j := range subArray {
			subArray[j] = false
		}
		flags[i] = subArray
	}
	ret := getNumberOfPoint(0, 0)
	fmt.Println(ret)
}
