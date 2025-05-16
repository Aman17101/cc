package main

import (
	"fmt"
	"time"
)

func main() {
	a()
	b()
	c()
	d()

	go a()
	go b()
	go c()
	go d()	
	time.Sleep(time.Second*3)

}

func a() {
	for i := 1; i <= 3; i++ {
		fmt.Println("function a is  :", i)
	}
}
func b() {
	for i := 1; i <= 3; i++ {
		fmt.Println("function b is :", i)
	}
}
func c() {
	for i := 1; i <= 3; i++ {
		fmt.Println("function c is  :", i)
	}
}
func d() {
	for i := 1; i <= 3; i++ {
		fmt.Println("function d is :", i)
	}
}
