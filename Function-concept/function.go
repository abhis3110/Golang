package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func ops(a int, b int) (int, int) {
	return a + b, a - b

}

func main() {
	result := sum(2, 2)
	fmt.Println(result)

	// function may return multiple values
	sum, subs := ops(2, 2)
	fmt.Println("2+2=", sum, "2-2=", subs)
	b, _ := ops(10, 2)
	fmt.Println("10+2=", b)
}
