package main

import "fmt"

func add() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
func main() {

	result := add()
	fmt.Println("result ", result())//1
	fmt.Println("result ", result())//2
	fmt.Println("result ", result())//3
	fmt.Println("result ", result())//4

}
