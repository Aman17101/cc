package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
	defer	fmt.Println(i)
	}

	//defer koi bhi gadbadi sambhalo
defer fmt.Println("hello")
defer fmt.Println("world")

fmt.Println("Aman")
defer fmt.Println("chef")
fmt.Println("Harry")


}
