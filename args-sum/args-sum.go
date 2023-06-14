package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithprogram := os.Args
	numA, err := strconv.Atoi(argsWithprogram[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	numB, err := strconv.Atoi(argsWithprogram[2])

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	num := numA + numB
	fmt.Println(num)
}
