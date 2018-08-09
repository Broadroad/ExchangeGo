package main

import (
	"flag"
	"strconv"
	"fmt"
)

var (
	x int
	y int
)

func init() {
	flag.IntVar(&x, "x", 40, "set x value")
	flag.IntVar(&y, "y", 40, "set y value")
}

// isGreatThan21 return true if (x, y) more than 21
func isGreatThan21(x, y int) bool {
	xstr := strconv.Itoa(x)
	ystr := strconv.Itoa(y)

	sum := 0
	for _, sx := range xstr {
		sum += int(sx - '0')
	}
	for _, sy := range ystr {
		sum += int(sy - '0')
	}

	if sum > 21 {
		return true
	}
	return false
}

func getNumberOfPoint() int{
	
}

func main() {
	flag.Parse()
	fmt.Println(getNumberOfPoint())
}
