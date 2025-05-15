package main

import (
	"fmt"

	"github.com/Aman17101/cc/importableExecutble/calculator"
	"github.com/Aman17101/cc/importableExecutble/store"
)

func main() {
	//	fmt.Println("sum =", calculator.Add(987872, 90909))
	fmt.Println("sub =", calculator.Sub(987872, 90909))
	fmt.Println("mul =", calculator.Mul(987872, 90909))
	fmt.Println("div =", calculator.Div(987872, 90909))

	usr := calculator.User{}
	usr.Name = "iuidai"
	fmt.Println("++++++", usr.Name)

//import using interface
	//var dboprs store.DBOperations
	dboprs :=store.Store{}
	dboprs.SaveRecord("just print this record ....")

}
