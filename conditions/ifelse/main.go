package main

import (
	//"flag"
	"fmt"
)

func main() {
	fmt.Println("conditions...")

	num:= 12
	if num <= 10 {
		fmt.Println("open dooor")
	} else if num == 11 {
		fmt.Println("ask then enter")
	} else {
		fmt.Println("not open the door")
	}


	numb :=14
	if (numb %2==0)&&(numb!=0){
		fmt.Println("even number")
	}else{
		fmt.Println("odd number")
	}


	event :="election"
	if event =="eligibilty criteria"{
		fmt.Println("above 18")
	}else if event =="may be"{
		fmt.Println(" only few days remain to complete 18")
	}else{ 
		fmt.Println("not eligible for",event)
	}
}
