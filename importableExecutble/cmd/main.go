package main

import (
	"fmt"
	"github.com/Aman17101/cc/importableExecutble/calculator"
)

func main() {
//	fmt.Println("sum =", calculator.Add(987872, 90909))
	fmt.Println("sub =", calculator.Sub(987872, 90909))
	fmt.Println("mul =", calculator.Mul(987872, 90909))
	fmt.Println("div =", calculator.Div(987872, 90909))

	usr:=calculator.User{}
	usr.Name="iuidai"
	fmt.Println("++++++",usr.Name)

}
