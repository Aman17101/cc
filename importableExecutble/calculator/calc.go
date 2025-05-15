package calculator

import "fmt"

type User struct{
	Name string
	age int
}

type student struct{
	Name string
	Rollno int
}

func add(a, b int) int {
	fmt.Println("add call hua")
	return a + b
}

//	func Sub( a ,b int)int{
//		return a - b
//	}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	return a / b
}