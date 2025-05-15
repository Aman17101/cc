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

type userOperation interface {
	addUser()
}

func main() {
	var aman UserA
	aman.name = "Aman"
	aman.address = "Delhi"
	aman.contactNo = 1234567890

	nikhil := UserB{
		name:      "Nikhil",
		address:   "mumbai",
		contactNo: 9876543210,
	}
	// aman.addUser()
	// nikhil.addUser()

	var user userOperation
	user = aman  //type UserA
	user.addUser()

	user =nikhil
	user.addUser()

}

func (aman UserA) addUser() {
	fmt.Println("about aman method ,", aman)

}
func (nikhil UserB) addUser() {
	fmt.Println("about nikhil method ,", nikhil)

}
