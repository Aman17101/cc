package main

import (
	"fmt"
	"os"
	
)

func main() {
	// i := []int{1, 6, 4, 3, 2}
	// fmt.Println(i)
	// fmt.Println(len(i))
	// fmt.Println(i[5])

	//fatal error(it stop the execution)
	dbconnection := false
	if !dbconnection {
		fatalerr("db connection failed :")
	}

	//user generated error

	pass := "12345666"

	fmt.Println(len(pass))
	if len(pass) < 8 {
		err := fmt.Errorf("password digit not match criteria :%v", pass)
		fmt.Println(err)
		//fatalerr("code criteria not matched \n")   	                                             //if you want to fatal error if criteria not matched 

		panic(err)
	                                            
	} else {
		savetodb(pass)
	}

}
func fatalerr(msg string) {
	fmt.Printf("some error  causes due to  fatal error occur %v\n", msg)
	os.Exit(1)
}

func savetodb(data  string) {
	fmt.Println("record saved success ,record ,",data)
}

