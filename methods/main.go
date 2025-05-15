package main

import "fmt"

type UserA struct {
	name      string
	address   string
	contactNo int
}

type UserB struct {
	name      string
	address   string
	contactNo int
}

func main() {

	var aman UserA
	aman.name = "AmanSingh"
	aman.address = "Delhi"
	aman.contactNo = 1234567890
	//aman.adduser(5)

	//var nikhil UserB
	nikhil := UserB{
		name:      "NikhilSingh",
		address:   "Mumbai",
		contactNo: 9876543210,
	}

	aman.adduser(5)
	fmt.Println("after adduser aman=", aman)
	aman.adduserdup(484748)
	fmt.Println("after adduserdup aman=", aman)

	nikhil.adduser(9)
	add(88966)
}

func add(a int) {
	fmt.Println("function chal raha , a =", a)
}

func (aman UserA) adduser(a int) {
	aman.name = "7233u2331247"
	fmt.Println("aman ka method chal raha hai ,aman", aman)
}
func (u *UserA) adduserdup(a int) int {

	u.name = "name value change kar raha hu"
	fmt.Println("aman ka  dup method chal raha hai ,aman", u)
	return 0
}
func (nikhil UserB) adduser(a int) {
	fmt.Println("nikhil ka method chal raha hai ,nikhil", nikhil)
}
