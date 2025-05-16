package main

import (
	"fmt"
	"os"
)

func main() {
	//defer recoverPanic()

	dbconnection := true

	 if !dbconnection {
		fatalErr("Database connection failed")
	}
	//normal case
	a := []string{
		"a wala string save hua",
	}
	savetodb(a)
	//panic case
	b := []string{}
	savetodb(b)
	//normal case
	c := []string{
		"c wala string save hua",
	}
	savetodb(c)
}

func recoverPanic() {
	r := recover()
	if r != nil {
		fmt.Println("kuch to gadbad hai :", r)
	}

}

func savetodb(record []string) {
	//defer recoverPanic()
	if len(record) < 1 {
		panicErr("no record found to save")
	} else {
		fmt.Println("saving to database", record)
	}
}
func fatalErr(msg interface{}) {
	fmt.Printf("fatal error cause due to pgm stop:%v ", msg)
	os.Exit(1)
}

func panicErr(data interface{}) {
	defer recoverPanic()
	fmt.Println("panic error occur: ")
	panic(data)
}
