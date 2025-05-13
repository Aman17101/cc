package main

import "fmt"

type Student struct {
	Name string
	Age int
	Grade int
	StudentAddress Address

}
type Address struct{
	lane1 string
	lane2 string
	po   string
	dist  string
}

func main() {

	var aman Student
	aman.Name = "Aman Singh"
	aman.Age = 20
	aman.Grade = 10

	shiva :=Student{
		Name:   "Shiva",
		//Rollno:  44,
		Age: 33,
		Grade: 10,
		StudentAddress :Address {
			lane1: "vns",
			lane2: "Gopiganj",
			po: "Pakhwaiya",
		dist :"mzp",
		
	},
}

	var amanaddress *Student
	amanaddress=&aman


	shivaddress := &shiva

	fmt.Println("aman address=" ,*amanaddress ,"shiva address =",shivaddress)

	
	
	fmt.Println("about pointer")
	

	fmt.Println(aman)
}