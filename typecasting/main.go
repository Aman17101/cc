package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	i = 10

	var a int16
	a = int16(i)

	fmt.Println("typecasting", a)

	var c string
	c = fmt.Sprintf("%d", a) //anything we have pass,it will convert into  string ( Sprintf())
	fmt.Println("c=", c)

	num := "1244457"
	int, error := strconv.Atoi(num)
	//strconv.Atoi() function convert string to integer
	fmt.Println("integer is :", int, "error is :", error)

	num = "aman"
	int, error = strconv.Atoi(num)
	//strconv.Atoi() function convert string to integer
	fmt.Println("integer is :", int, "error is :", error)

}
