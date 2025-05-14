package main

import "fmt"

func main() {

	for i := 100; i >0; i--{
		fmt.Println(i)

	}

	//2 example
	var i int = 1
	for {
		if i > 10 {
			break
		}
		i++
		fmt.Println("i++++", i)
	}

	//third example

	i = 10
	condition := true
	for condition {
		if i >100 {
			condition = false
		} else {
			i++
			fmt.Println("this is different loop :", i)

		}
	}

	//range

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
		

// for i:=0;i<len(arr);i++{
// 	fmt.Println("array list",arr[i])
// }
var singh []string
 singh = append(singh,("AMan"))
 singh = append(singh,("rorhan"))
 singh = append(singh,("Ammmm"))
 singh = append(singh,("nnahsj"))
//fmt.Println("arr of student", singh[0])
for i:=0;i<len(singh);i++{
	fmt.Println("array list",singh[i])
}

for _ ,j :=range arr{
	fmt.Println("array list",j)
}


	//map declaration
	m := make(map[string]int)
	//save data in map
	m["one"] = 1
	m["pqr"] = 10
	m["asa"] = 100
	m["jkl"] = 1000
	
	


	//read data from maps

	for key,val:=range m{
		fmt.Println("key",key,"value",val)
	}
}

