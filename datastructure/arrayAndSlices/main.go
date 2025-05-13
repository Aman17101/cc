package main

import "fmt"

func main() {
	

	var arr [10]int
		arr[0]=10
		arr[1]=20
		arr[2]=30
	arr[3]=100
		arr[4]=50
		arr[5]=68
		arr[6]=34
		arr[7]=28
		arr[9]=77
arr[8]=8
		
fmt.Println("array list",arr)

arr[3]=40


//string of array
var names [10]string
		names [0]="10"
		names [1]="20"
		names [2]="30"
	names [3]= "100"
		names [4]="50"
		names [5]="68"
		names [6]="34"
		names [7]="28"
		names [9]="77"
names [8]="8"
		
fmt.Println("array list",names )

names [3]="40"
  


//slices
var Student []string
Student = append(Student, "Aman")
Student =append(Student, "Rahul")
fmt.Println("student",Student[0])
fmt.Println("student",Student[1])


	
	

}
