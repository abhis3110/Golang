package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var whatToSay string // variable declaration of type string
	whatToSay = "Goodbye"
	println(whatToSay)

	var i int = 7
	fmt.Println("The value of i set to", i)

	// automatic assignment of type of variable
	whatWasSaid, whatHaveSaid := saySomething()
	fmt.Println("The function return", whatWasSaid, whatHaveSaid)
}

func saySomething() (string, string) {
	// function can return more than two variable in Golang
	return "something", "else"
}
