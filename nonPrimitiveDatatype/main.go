package main

import "fmt"

//non primitive datatype

type Student struct {
	Name   string
	Rollno int
	Age    int
	StudentAddress Address
}
type Address struct {
		lane1 string
		dist string
		post string
		vill string
	}
	// type phone struct{
	// 	model string
	// 	ram string
	// 	imei int
	// 	Config
	// }

	// type laptop struct{
	// 	brand string
	// 	model string
	// 	ram string
	// 	configuration config
	// }
	// type Config struct{
	// 	OS string
	// 	processor string
	// 	ram string
	// }


func main() {

	var aman Student
	//var nick Student

	aman.Name="Aman singh"
	aman.Age =22
	aman.StudentAddress.dist="Mirzapur"

	nick :=Student{
		Name:   "Nick",
		Rollno:  44,
		Age: 33,
		StudentAddress :Address {
		dist :"mzp",
		vill :"pap",
		post :"pzp",
	},
	}

	//interface 
	val :=1222
	val1 :="8777"

	var interfaceExample interface{}

	interfaceExample = val
	fmt.Println("interfacevalue is=",interfaceExample)

	interfaceExample = val1
	fmt.Println("interfacevalue is= ",interfaceExample)

	interfaceExample = false
	fmt.Println("interfacevalue is =",interfaceExample)

	

	// fmt.Println("hello team ....", aman)
	
	
fmt.Println("team address is ",aman)
 fmt.Println("hello team ....nick struct", nick)
	
}
