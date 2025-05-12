package main

import "fmt"

func main() {
	var number int
	number = 22
	var  addressofNumber *int 
	addressofNumber =&number

	number3 :=223434
fmt.Println(number3)

num4 :=number
fmt.Println(  "number .....",num4)

	var decimalNumber float32
	decimalNumber = 22.5
	var addressofdecimalNumber *float32
	 addressofdecimalNumber =&decimalNumber	

	var flags bool
	flags = true
	var addressofflags *bool
	addressofflags =&flags

	var name string
	name = "John"
	var addressofname *string
	addressofname =&name

	fmt.Printf("number is = %d, decimal number is =%f,flags value is =%v, name value is =%v ", number, decimalNumber, flags, name)
	fmt.Printf("address of number is = %v,address of  decimal number is =%v,address of flags value is =%v, address of name value is =%v ",addressofNumber,addressofdecimalNumber,addressofflags ,addressofname )
}
