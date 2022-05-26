package main

import "fmt"

type employee struct {
	fname, lname string
}

func (e employee) getFullName() string {
	//e.lname = "Sharma"
	return e.fname + " " + e.lname
}

func (e *employee) changeName() {
	e.lname = "Sharma"
}

func main() {

	e := employee{"Rohit", "Sangamnerkar"}
	fmt.Println(e)
	fmt.Println("Full Name : ", e.getFullName())
	e.changeName()
	fmt.Println(e)

}
