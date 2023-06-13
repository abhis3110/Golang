package main

import (
	"fmt"
	"os"
)

func main() {
	argswithprogram := os.Args
	argswithoutprogram := os.Args[1:]
	args := os.Args[3]
	fmt.Println(argswithprogram)
	fmt.Println(argswithoutprogram)
	fmt.Println(args)

}
