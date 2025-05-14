package main

import (
	"fmt"
)

func main() {
	//map declaration
	m := make(map[string]int)
	//save data in map
	m["one"] = 1
	m["pqr"] = 10

	//read data from maps
	llj, kk:= m["pqr"]
fmt.Println("val=", llj, "ok=", kk)

val ,ok :=m["abc"]
if ok{
	fmt.Println("val found ,val=", val)
}else{
	fmt.Println("val not found")
}

	//fmt.Println("val=", val, "ok=", ok)


}
