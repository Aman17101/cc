package main

import "fmt"

func main() {
	f :=5
	g := 6
        sum :=add(f,g)
	fmt.Println("add of a and b :",sum)
	fmt.Println("sub of a and b :", sub(f, g))
	fmt.Println("mul of a and b :", mul(f, g))

	//for failing division
	divide,err :=div(f,0)
	if err ==nil{
		fmt.Println("div of a and b :", divide)
	}else{
	fmt.Println(" divide hi nhi hua ,err ki val h :", err)
	}

	//for passing division
	divide,err =div(f,2)
	if err ==nil{
		fmt.Println("div of a and b :", divide)
	}else{
	fmt.Println(" divide hi nhi hua ,err ki val h :", err)
	}

	fmt.Println("f=",f,"g=",g)
}

//for adding two number

func add(a, b int) int {
	return a + b
}

//for subtracting two number

func sub(a, b int) int {
	return a - b
}

//for multiply of two number

func mul(a, b int) int {
	return a * b
}

//for dividing two number

func div(d, e int) (int, error) {
	if e != 0 {
		return d / e, nil

	} else {
		return 0, fmt.Errorf("dividing with zero  not possible")
	}

}
