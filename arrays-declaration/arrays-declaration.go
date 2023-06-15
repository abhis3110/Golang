package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println(b)

	c := [5]int{0, 1, 2}
	fmt.Println(c)
	fmt.Println(c[2])

	for i := 0; i < 5; i++ {
		fmt.Print(b[i])
	}

}
