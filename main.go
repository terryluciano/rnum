package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func getRandomNumer(start int, end int) int {

    if start == end {
        return start
    }

    if start > end {
        start, end = end, start
    }

    return (rand.Intn(end - start + 1) + start)

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
        if err != nil {
            fmt.Println("Invalid start value")
            os.Exit(1)
        }

        result = getRandomNumer(1, end)

        fmt.Println(result)
    } else {
        start, err = strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid start value")
        }

        end, err = strconv.Atoi(args[1])
        if err != nil {
            fmt.Println("Invalid end value")
        }

        result = getRandomNumer(start, end)
        
        fmt.Println(result)
    }

    os.Exit(0)
}