package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func getRandomNumber(start int, end int) int {
	if start == end {
		return start
	}

	if start > end {
		start, end = end, start
	}

	return (rand.Intn(end-start+1) + start)
}

func handleError(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
		os.Exit(1)
	}
}

func main() {
	var start int
	var end int
	var result int
	var err error

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		end, err = strconv.Atoi(args[0])
		handleError(err, "Invalid end value")

		result = getRandomNumber(1, end)

		fmt.Println(result)
	} else {
		start, err = strconv.Atoi(args[0])
		handleError(err, "Invalid start value")

		end, err = strconv.Atoi(args[1])
		handleError(err, "Invalid end value")

		result = getRandomNumber(start, end)

		fmt.Println(result)
	}

	os.Exit(0)
}
