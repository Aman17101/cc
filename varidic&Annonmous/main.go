package main

import "fmt"

func main() {

	func  (a int)int{
		fmt.Println("annomous function is running " )
		return 0
	}(7)


	fmt.Println("sum of a and b is =",add(5,6,7,1))
}
func add( b... int) int {

	sum:=0
	for _,val := range b{
	sum +=val
}
return sum 
	
}  